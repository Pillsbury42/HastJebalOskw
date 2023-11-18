// Server/Node
package main

import (
	"context"
	"fmt"
	"time"
	// "io"
	"log"
	"net"

	// this has to be the same as the go.mod module,
	// followed by the path to the folder the proto file is in.
	gRPC "github.com/Pillsbury42/HastJebalOskw/handin5/gRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

var isLeader = false
var nodeList []Node //List of nodes, ideally in order of priority
var bidderIDs map[int64]struct{} //golang doesn't have sets, but a map with empty struct as value has the same effect
var topBid struct {
	bidderID   int64
	highestBid int64
}
var myID int64

type Node struct {
	nodeClient   gRPC.NewAuctionClient
	nodeId int64
}
var endTime time.Time //The ending time of the auction

func main() {
	//f := setLog()
	//defer f.Close()

	fmt.Println(".:Node is starting:.")

	go launchServer(myport)
	//This must be able to connect to multiple other nodes
	server:=ConnectToServer(otherport)

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

	gRPC.RegisterAuctionServer(grpcServer, gRPC.UnimplementedAuctionServer)

	log.Printf("Listening at %v\n\n", list.Addr())

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	// code here is unreachable because grpcServer.Serve occupies the current thread.
}
func ConnectToServer(port string) gRPC.NewAuctionClient {
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
func (s *gRPC.UnimplementedAuctionServer) Bid(ctx context.Context, msg *gRPC.BidMessage) (*gRPC.BidReplyMessage, error) {
	//The node receives a bid from a client, processes this, copies it to the others, and returns a bid reply message,
	
	//Process:
	if (isLeader){
		//Has this client bid before?
		if _, ok := bidderIDs[msg.bidderID]; !ok {
			bidderIDs[msg.bidderID] =struct{}{}
		}
		//Is this bid higher than the current max?
		if topBid.highestBid<msg.bidamount {
			topBid.bidderID=msg.bidderID
			topBid.highestBid=msg.highestBid
			for _, element := range nodeList {
				element.nodeClient.Bidupdate(context.Background(), msg)
			}
		}
	} else { //If the node is not the leader, then it means the leader must have crashed
		for _, element := range nodeList {
			if (element.nodeId > myID){
				electmsg := gRPC.ElectionMessage{
					
				}
				element.nodeClient.Election(context.Background(), electmsg)
			}
		}
	}
}
func (s *gRPC.UnimplementedAuctionServer) BidUpdate(ctx context.Context, msg *gRPC.BidUpdateMessage) (*gRPC.EmptyMessage, error) {
	//The node receives a message from the leader telling it what the new highest bid is and updates accordingly
}
func (s *gRPC.UnimplementedAuctionServer) Result(ctx context.Context, msg *gRPC.ResultMessage) (*gRPC.ResultReplyMessage, error) {
	//The node receives a query for the current status of the auction from a client,
	//and returns a reply either detailing who won the auction, or what the current highest bid is and how long is left
	//Once again, if this node is not the leader, then an election is called
}
func (s *gRPC.UnimplementedAuctionServer) Election(ctx context.Context, msg *gRPC.ElectionMessage) (*gRPC.ElectionReplyMessage, error) {
	//The node receives a query from another node, asking who the highest alive node is
	//It then iterates through all nodes higher than itself, asking them the same question
	//It then takes the answer from the highest node and returns it in an election reply message
	//If it is the highest node alive, i.e. if none of the other nodes answer, 
	//then it sends out coordinator messages to all other nodes telling them that it is the new leader
}
func (s *gRPC.UnimplementedAuctionServer) Coordinator(ctx context.Context, msg *gRPC.CoordinatorMessage) (*gRPC.EmptyMessage, error) {
	//The node receives a message from another node, telling it that the other node is the new leader
	//The node updates its leader-variable accordingly.
}
