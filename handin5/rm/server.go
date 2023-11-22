// Server/Node
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	// this has to be the same as the go.mod module,
	// followed by the path to the folder the proto file is in.
	gRPC "github.com/Pillsbury42/HastJebalOskw/handin5/gRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var isLeader = false
var nodeList []Node                 //List of nodes, ideally in order of priority
var bidderNames map[string]struct{} //golang doesn't have sets, but a map with empty struct as value has the same effect
var myID int64
var topBid struct {
	bidderName string
	highestBid int64
}

var myport = flag.String("port", "5400", "receiving port")

// This is just a wrapper to inherit
type ImplementedAuctionServer struct {
	gRPC.UnimplementedAuctionServer
}
type Node struct {
	nodeClient gRPC.AuctionClient
	nodeId     int64
}

var startTime int64 //The start time of the auction, which is compared against when determining if the auction is over

func main() {
	//f := setLog()
	//defer f.Close()

	startTime = 0

	fmt.Println(".:Node is starting:.")
	flag.Parse()
	go launchServer(*myport)
	//This must be able to connect to multiple other nodes
	for _, element := range []string{"5400", "5401", "5402"} {
		if element != *myport {
			nodeList = append(nodeList, ConnectToServer(element))
		}
	}

	for true {
		//an infinite loop to ensure the program keeps running
	}
}

func launchServer(port string) {
	log.Printf("Attempting to create listener on port %s\n", port)

	// Create listener tcp on given port or default port 5400
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Printf("Failed to listen on port %s: %v", port, err)
		return
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	gRPC.RegisterAuctionServer(grpcServer, ImplementedAuctionServer{})

	log.Printf("Listening at %v\n\n", list.Addr())

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	// code here is unreachable because grpcServer.Serve occupies the current thread.
}

func ConnectToServer(port string) gRPC.AuctionClient {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	log.Printf("Attempts to dial on port %s\n", port)
	conn, err := grpc.Dial(fmt.Sprintf(":%s", port), opts...)
	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return nil
	}
	server := gRPC.NewAuctionClient(conn)
	log.Println("the connection is: ", conn.GetState().String())
	return server
}

func (s ImplementedAuctionServer) Bid(ctx context.Context, msg *gRPC.BidMessage) (*gRPC.BidReplyMessage, error) {
	//The node receives a bid from a client, processes this, copies it to the others, and returns a bid reply message,
	var reply *gRPC.BidReplyMessage
	//Checks if the timer has started. If not, starts the timer and shares the start time with the other nodes
	if startTime == 0 {
		startTime = time.Now().Unix()
		timeMsg := &gRPC.StartMessage{
			StartTime: startTime,
		}
		for _, element := range nodeList {
			element.nodeClient.Start(context.Background(), timeMsg)
		}
	}
	//Process:
	if time.Now().Unix()-startTime < 100 {
		if isLeader {
			//Has this client bid before?
			if _, ok := bidderNames[msg.BidderName]; !ok {
				bidderNames[msg.BidderName] = struct{}{}
			}
			//Is this bid higher than the current max?
			if topBid.highestBid < msg.Amount {
				topBid.bidderName = msg.BidderName
				topBid.highestBid = msg.Amount
				for _, element := range nodeList {
					element.nodeClient.Bidupdate(context.Background(), msg)
				}
			}
			reply = &gRPC.BidReplyMessage{
				Success: true,
			}
		} else { //If the node is not the leader, then it means the leader must have crashed
			for _, element := range nodeList {
				if element.nodeId > myID {
					element.nodeClient.Election(context.Background(), &gRPC.EmptyMessage{})
				}
			}
			reply = &gRPC.BidReplyMessage{
				Success: false,
			}
		}
	} else {
		reply = &gRPC.BidReplyMessage{
			Success: false,
		}
	}

	return reply, nil

}

func (s ImplementedAuctionServer) BidUpdate(ctx context.Context, msg *gRPC.BidMessage) (*gRPC.EmptyMessage, error) {
	//The node receives a message from the leader telling it what the new highest bid is and updates accordingly
	topBid.bidderName = msg.BidderName
	topBid.highestBid = msg.Amount
	return &gRPC.EmptyMessage{}, nil
}
func (s ImplementedAuctionServer) Result(ctx context.Context, msg *gRPC.EmptyMessage) (*gRPC.ResultReplyMessage, error) {
	//The node receives a query for the current status of the auction from a client,
	//and returns a reply either detailing who won the auction, or what the current highest bid is and how long is left
	//Once again, if this node is not the leader, then an election is called
	if time.Now().Unix()-startTime < 100 {
		if isLeader {
			return &gRPC.ResultReplyMessage{Over: false, WinnerName: topBid.bidderName, Highest: topBid.highestBid}, nil
		} else {
			for _, element := range nodeList {
				if element.nodeId > myID {
					element.nodeClient.Election(context.Background(), &gRPC.EmptyMessage{})
				}
			}
			return &gRPC.ResultReplyMessage{Over: false, WinnerName: topBid.bidderName, Highest: topBid.highestBid}, nil
		}
	} else {
		return &gRPC.ResultReplyMessage{Over: false, WinnerName: topBid.bidderName, Highest: topBid.highestBid}, nil
	}
}

func (s ImplementedAuctionServer) Election(ctx context.Context, msg *gRPC.EmptyMessage) (*gRPC.ElectionReplyMessage, error) {

	//The node receives a query from another node, asking who the highest alive node is
	//It then iterates through all nodes higher than itself, asking them the same question
	//It then takes the answer from the highest node and returns it in an election reply message
	//If it is the highest node alive, i.e. if none of the other nodes answer,
	//then it sends out coordinator messages to all other nodes telling them that it is the new leader
	var countHigher = 0
	var highestID = myID
	for _, node := range nodeList {
		if node.nodeId > myID {
			countHigher++
			highestID = node.nodeId
			node.nodeClient.Election(context.Background(), &gRPC.EmptyMessage{})
		}
	}
	//If no other nodes respond, then I am the winner
	if countHigher == 0 {
		coordmsg := &gRPC.CoordinatorMessage{
			CoordID: myID,
		}
		for _, node := range nodeList {
			node.nodeClient.Coordinator(context.Background(), coordmsg)
		}
	}

	reply := &gRPC.ElectionReplyMessage{
		ReplyID: highestID,
	}
	return reply, nil

}

func (s ImplementedAuctionServer) Coordinator(ctx context.Context, msg *gRPC.CoordinatorMessage) (*gRPC.EmptyMessage, error) {
	//The node receives a message from another node, telling it that the other node is the new leader
	//The node updates its leader-variable accordingly.
	isLeader = false
	return &gRPC.EmptyMessage{}, nil
}

func (s ImplementedAuctionServer) Start(ctx context.Context, msg *gRPC.StartMessage) (*gRPC.EmptyMessage, error) {
	//The node receives a message from another node, telling it that the auction has started and what time it started

	startTime = msg.startTime
	return &gRPC.EmptyMessage{}, nil
}
