package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

gRPC "github.com/Pillsbury42/HastJebalOskw/handin4/gRPC"

"google.golang.org/grpc"
"google.golang.org/grpc/credentials/insecure"
// Same principle as in client. Flags allows for user specific arguments/values
var myName = flag.String("name", "default", "Senders name")
var myPort = flag.String("server", "5400", "Tcp server")

var myNode gRPC.MutexClient   //the server
var myConn *grpc.ClientConn //the server connection

type Node struct {
	gRPC.UnimplementedMutexServer
	name string
	port string
	nextport string
	id int
	voted bool
}

func main() {
	f := setLog() //uncomment this line to log to a log.txt file instead of the console
	defer f.Close()

	// This parses the flags and sets the correct/given corresponding values.
	flag.Parse()
	fmt.Println(".:node is starting:.")
	launchNode()
	ConnectToNode()
}

func ConnectToNode() {}

func launchNode() {
	//launch
	log.Printf("Node %s: Attempts to create listener on port %s\n", *myName, *myPort)
	//create listener
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *myPort))
	if err != nil {
		log.Printf("Server %s: Failed to listen on port %s: %v", *myName, *myPort, err)
		return
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// makes a new server instance using the name and port from the flags.
	server := &Node{
		name:        *serverName,
		port:        *port,
		
	}
	gRPC.RegisterMutexServer(grpcServer,server)
	log.Printf("Server %s: Listening at %v\n\n", server.name, list.Addr())

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}

func (s *Node) Election(ctx context.Context, elmsg ElectionMessage) (*EmptyMessage, error) {
	voteID:=s.id
	if (elmsg.topcnID > id) {
		voteID=elmsg.topcnID}
	if !s.voted {
		s.voted=true
		
			msg := &gRPC.ElectionMessage {
				topcnID: voteID
			}
			ack, _ := client.Election(context.Background(), msg)
		
	} else {
			coordmsg := &gRPC.CoordinatorMessage {
				coordID = voteID
			}
			ack, _ := client.Elected(context.Background(), coordmsg)
		
		}
	
	empty := &gRPC.EmptyMessage{}
	return (empty, n)
}

func (s *Node) Elected(ctx context.Context, coordmsg CoordinatorMessage) (*EmptyMessage, error) {
	
	empty := &gRPC.EmptyMessage{}
	return (empty, nil)
}



