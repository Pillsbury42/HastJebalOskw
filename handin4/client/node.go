package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	gRPC "github.com/Pillsbury42/HastJebalOskw/handin4/gRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//Clients send Election and Coordination messages to the server, which sends empty messages back. Therefore the server equals the next node in line

// Same principle as in client. Flags allows for user specific arguments/values
var myName = flag.String("name", "default", "This node's name")
var listenPort = flag.String("clientp", "5400", "receiving port")
var serverPort = flag.String("serverp", "5401", "sending port")
var hasToken = flag.String("hastoken", "false", "Starts with token")
func (s *gRPC.UnimplementedAuctionServer) Bid(ctx context.Context, msg *gRPC.BidMessage) (*gRPC.BidReplyMessage, error) {
	//The node receives a bid from a client, processes this, copies it to the others, and returns a bid reply message,
	//If the node is not the leader, then it means the leader must have crashed, so it calls an election
}
var nextNode gRPC.MutexClient //the server
var nextConn *grpc.ClientConn //the "server" connection, used to check if the other node is responding and to close the connection

// The node struct is needed to handle
type Node struct {
	gRPC.UnimplementedMutexServer
	name        string
	port        string
	nextport    string
	wantsAccess bool
}

func main() {
	flag.Parse()
	f := setLog() //uncomment this line to log to a log.txt file instead of the console
	defer f.Close()

	// This parses the flags and sets the correct/given corresponding values.

	fmt.Println(".:node is starting:.")

	go launchNode()

	//send the first message if initialized with token
	if *hasToken == "true" {
		isConnected = true
		connectToNode()
		log.Printf("Starting the token pass\n")
		nextNode.HasToken(context.Background(), &gRPC.HasTokenMessage{Token: true})
	} else {
		log.Printf("I did not start with the token\n")
	}
	for true {
	}
}

func launchNode() {
	//This is equivalent o the launchServer() method
	//This creates a listener at this node's port and sets local a Node struct
	//All nodes are servers and clients

	//launch
	log.Printf("Node %s: Attempts to create listener on port %s\n", *myName, *listenPort)
	//create listener
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *listenPort))
	if err != nil {
		log.Printf("Server %s: Failed to listen on port %s: %v", *myName, *listenPort, err)
		return
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// makes a new server instance using the name and port from the flags.
	meNode := &Node{
		name:        *myName,
		port:        *listenPort,
		nextport:    *serverPort,
		wantsAccess: (rand.Intn(2) < 1),
	}
	gRPC.RegisterMutexServer(grpcServer, meNode)
	log.Printf("Server %s: Listening at %v\n\n", meNode.name, list.Addr())
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}

func connectToNode() {
	//This is equivalent to client's ConnectToServer method
	//Here, our node dials up the connection to the next node in the ring
	//If succesful, this sets the next node as "server"
	//dial options
	//the server is not using TLS, so we use insecure credentials
	//(should be fine for local testing but not in the real world)
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	//dial the server, with the flag "server", to get a connection to it
	log.Printf("Node %s: Attempts to dial on port %s\n", *myName, *serverPort)
	conn, err := grpc.Dial(fmt.Sprintf(":%s", *serverPort), opts...)
	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return
	}

	// makes a client from the server connection and saves the connection
	// and prints whether or not the connection was is READY
	nextNode = gRPC.NewMutexClient(conn)
	nextConn = conn
	log.Printf("the connection is: %s \n", conn.GetState().String())
}

func (s *Node) HasToken(ctx context.Context, msg *gRPC.HasTokenMessage) (*gRPC.EmptyMessage, error) {
	if !isConnected {
		connectToNode()
		isConnected = true
	}
	//if the node has the token, it can access the critical section. Otherwise the token is passed on
	if s.wantsAccess && msg.Token {
		log.Printf("Node %s has token and wants access. Accessing critical section...\n", s.name)
		time.Sleep(2 * time.Second)
		s.wantsAccess = false
		log.Printf("Node %s has accessed the critical section\n", s.name)
	} else {
		log.Printf("Node %s has token, but doesn't want access.\n", s.name)

	}
	log.Printf("Node %s has finished and is passing the token on\n", s.name)

	if rand.Intn(2) < 1 {
		log.Printf("Node %s wants access\n", s.name)
		s.wantsAccess = true
	} else {
		log.Printf("Node %s does not want access\n", s.name)
	}

	nextNode.HasToken(context.Background(), msg)

	return &gRPC.EmptyMessage{}, nil
}

// sets the logger to use a log.txt file instead of the console
func setLog() *os.File {
	// Clears the log.txt file when a new server is started
	if err := os.Truncate("log_"+*myName+".txt", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	// This connects to the log file/changes the output of the log informaiton to the log.txt file.
	f, err := os.OpenFile("log_"+*myName+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return f
}
