// Server/Node
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
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
var leader Node

var myport = flag.String("port", "5400", "receiving port")
var leaderPort = flag.String("leader", "5402", "leader's port")

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
	f := setLog()
	defer f.Close()

	bidderNames = make(map[string]struct{})

	startTime = 0

	fmt.Println(".:Node is starting:.")
	flag.Parse()

	if *myport == *leaderPort {
		isLeader = true
		fmt.Println("I am leader")
	}

	if *myport == "5400" {
		myID = 1
	} else if *myport == "5401" {
		myID = 2
	} else if *myport == "5402" {
		myID = 3
	}
	go launchServer(*myport)
	//This must be able to connect to multiple other nodes
	for _, element := range []string{"5400", "5401", "5402"} {
		if element != *myport {
			othernode := ConnectToServer(element)
			nodeList = append(nodeList, othernode)
			leader = Node{}
			if element == *leaderPort {
				leader = othernode
			}
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

func ConnectToServer(port string) Node {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	log.Printf("Attempts to dial on port %s\n", port)
	conn, err := grpc.Dial(fmt.Sprintf(":%s", port), opts...)
	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return Node{nil, myID}
	}
	server := gRPC.NewAuctionClient(conn)
	log.Println("the connection is: ", conn.GetState().String())
	return Node{server, myID}
}

func (s ImplementedAuctionServer) Bid(ctx context.Context, msg *gRPC.BidMessage) (*gRPC.BidReplyMessage, error) {
	//The node receives a bid from a client, processes this, copies it to the others, and returns a bid reply message,
	var reply *gRPC.BidReplyMessage
	//Checks if the timer has started. If not, starts the timer and shares the start time with the other nodes
	if startTime == 0 {
		startTime = time.Now().Unix()
		log.Println("Starting auction at Unix-time %d", startTime)
		timeMsg := &gRPC.StartMessage{
			StartTime: startTime,
		}
		for _, element := range nodeList {
			element.nodeClient.Start(context.Background(), timeMsg)
		}
	}
	//Process:
	timeleft := time.Now().Unix() - startTime
	if timeleft < 100 {
		log.Println("Seconds left until end of auction: %d", timeleft)
		if isLeader {
			//Has this client bid before?
			if _, ok := bidderNames[msg.BidderName]; !ok {
				log.Println("%s is a new client, adding to list of bidders", msg.BidderName)
				bidderNames[msg.BidderName] = struct{}{}
			} else {
				log.Println("%s has bidded before", msg.BidderName)
			}
			//Is this bid higher than the current max?
			if topBid.highestBid < msg.Amount {
				log.Println("%d is a larger bid than the current highest of %d", msg.Amount, topBid.highestBid)
				topBid.bidderName = msg.BidderName
				topBid.highestBid = msg.Amount
				for _, element := range nodeList {
					element.nodeClient.Bidupdate(context.Background(), msg)
				}
				reply = &gRPC.BidReplyMessage{
					Success:  "Success",
					LeaderID: myID,
				}
			} else {
				log.Println("%d is not a larger bid than the current highest of %d", msg.Amount, topBid.highestBid)
				reply = &gRPC.BidReplyMessage{
					Success:  "LowBid",
					LeaderID: myID,
				}
			}

		} else {
			//If the node is not the leader, then it means either that:
			// the client is not updated on who is the leader, or the leader must have crashed
			// First, ask the leader yourself
			ack, err := leader.nodeClient.Bid(context.Background(), msg)
			if err != nil {
				log.Println("No response from leader, calling an election")
				//timeout code here, ie. election
				elect()
				//time.Sleep(1 * time.Second)

				reply = &gRPC.BidReplyMessage{
					Success:  "Election",
					LeaderID: leader.nodeId,
				}
				return reply, nil

			}
			//this node is not the leader, but a leader exists
			return ack, err
		}
	} else {
		log.Println("Auction ended with the winner %s, bidding %d", topBid.bidderName, topBid.highestBid)
		reply = &gRPC.BidReplyMessage{
			Success:  "Over",
			LeaderID: leader.nodeId,
		}
	}

	return reply, nil

}

func (s ImplementedAuctionServer) Bidupdate(ctx context.Context, msg *gRPC.BidMessage) (*gRPC.EmptyMessage, error) {
	//The node receives a message from the leader telling it what the new highest bid is and updates accordingly
	if _, ok := bidderNames[msg.BidderName]; !ok {
		bidderNames[msg.BidderName] = struct{}{}
	}
	log.Println("Updating top bidder name %s and amount %d", topBid.bidderName, topBid.highestBid)
	topBid.bidderName = msg.BidderName
	topBid.highestBid = msg.Amount
	return &gRPC.EmptyMessage{}, nil
}
func (s ImplementedAuctionServer) Result(ctx context.Context, msg *gRPC.EmptyMessage) (*gRPC.ResultReplyMessage, error) {
	//The node receives a query for the current status of the auction from a client,
	//and returns a reply either detailing who won the auction, or what the current highest bid is and how long is left
	//Once again, if this node is not the leader, then an election is called

	if isLeader {
		if time.Now().Unix()-startTime < 100 {
			log.Println("Auction is NOT over. The highest bid so far is %d by bidder %s", topBid.highestBid, topBid.bidderName)
			return &gRPC.ResultReplyMessage{WinnerName: topBid.bidderName, Highest: topBid.highestBid, LeaderID: myID, Success: "NotOver"}, nil
		} else {
			log.Println("Auction is over. The highest bid so far is %d by bidder %s", topBid.highestBid, topBid.bidderName)
			return &gRPC.ResultReplyMessage{WinnerName: topBid.bidderName, Highest: topBid.highestBid, LeaderID: leader.nodeId, Success: "Over"}, nil
		}

	} else {
		ack, err := leader.nodeClient.Result(context.Background(), msg)
		if err != nil {
			log.Println("No response from leader, calling an election")
			//timeout code here, ie. election
			elect()
			//time.Sleep(1 * time.Second)
			return &gRPC.ResultReplyMessage{LeaderID: leader.nodeId, Success: "Election"}, err

		}
		return ack, err
	}
}

func (s ImplementedAuctionServer) Election(ctx context.Context, msg *gRPC.EmptyMessage) (*gRPC.ElectionReplyMessage, error) {

	return elect(), nil

}
func elect() *gRPC.ElectionReplyMessage {
	//The node receives a query from another node, asking who the highest alive node is
	//It then iterates through all nodes higher than itself, asking them the same question
	//It then takes the answer from the highest node and returns it in an election reply message
	//If it is the highest node alive, i.e. if none of the other nodes answer,
	//then it sends out coordinator messages to all other nodes telling them that it is the new leader
	var countHigher = 0
	var highestID = myID
	for _, node := range nodeList {
		if node.nodeId > myID {
			_, err := node.nodeClient.Election(context.Background(), &gRPC.EmptyMessage{})
			if err == nil {
				countHigher++
				highestID = node.nodeId

			}
		}
	}
	//time.Sleep(200 * time.Millisecond)
	//If no other nodes respond, then I am the winner
	if countHigher == 0 {
		log.Println("I am the node with the biggest ID alive, I am the new leader")
		isLeader = true
		leader.nodeId = myID
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
	return reply
}

func (s ImplementedAuctionServer) Coordinator(ctx context.Context, msg *gRPC.CoordinatorMessage) (*gRPC.EmptyMessage, error) {
	//The node receives a message from another node, telling it that the other node is the new leader
	//The node updates its leader-variable accordingly.
	log.Println("Node %d has the biggest ID alive and is the new leader", msg.CoordID)
	isLeader = false
	leader.nodeId = msg.CoordID
	for _, node := range nodeList {
		if node.nodeId == msg.CoordID {
			leader.nodeClient = node.nodeClient
			break
		}
	}
	return &gRPC.EmptyMessage{}, nil
}

func (s ImplementedAuctionServer) Start(ctx context.Context, msg *gRPC.StartMessage) (*gRPC.EmptyMessage, error) {
	//The node receives a message from another node, telling it that the auction has started and what time it started

	log.Println("Node %d is starting auction at Unix-time %d", leader.nodeId, startTime)
	startTime = msg.StartTime
	return &gRPC.EmptyMessage{}, nil
}

func setLog() *os.File {
	f, err := os.OpenFile("log_"+*myport+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return f
}
