package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	gRPC "github.com/Pillsbury42/HastJebalOskw/handin4/gRPC"
)
//Clients send Election and Coordination messages to the server, which sends empty messages back. Therefore the server equals the next node in line

// Same principle as in client. Flags allows for user specific arguments/values
var myName = flag.String("name", "default", "Senders name")
var listenPort = flag.String("clientp", "default", "Tcp server")
var serverPort = flag.String("serverp", "default", "Server port")

<<<<<<< HEAD
var nextNode gRPC.MutexClient   //the server
=======
var myNode gRPC.MutexClient   //the server
>>>>>>> d0e6741bf395f59cc902b67abdaca541da6ecec5
var myConn *grpc.ClientConn //the "server" connection, used to check if the other node is responding and to close the connection
// The node struct is needed to handle
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

	//write -clientp <port>  as cmd line arg
	launchNode()

	//write -serverp <port> as cmd line arg
	connectToNode()
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
		
	}
	gRPC.RegisterMutexServer(grpcServer,meNode)
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
	log.Printf("Node %s: Attempts to dial on port %s\n", *myName, serverPort)
	conn, err := grpc.Dial(fmt.Sprintf(":%s", *serverPort), opts...)
	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return
	}

	// makes a client from the server connection and saves the connection
	// and prints whether or not the connection was is READY
	net = gRPC.NewMutexClient(conn)
	ServerConn = conn
	log.Println("the connection is: ", conn.GetState().String())
}

func (s *Node) Election(ctx context.Context, elmsg gRPC.ElectionMessage) (*gRPC.EmptyMessage, error) {
	//id is a rank/id 
	//If voted then we've completed a lap and must start coordination, otherwise:
	//set voted to true. If own id is higher than message id, send message with own id, otherwise pass on
	voteID:=s.id
	
	if (elmsg.topcnID > id) {
		voteID=elmsg.topcnID
	}
	
	if !s.voted {
		s.voted=true
		
			msg := &gRPC.ElectionMessage {
				topcnID: voteID,
			}
			ack, _ := client.Election(context.Background(), msg)
		
	} else {
			coordmsg := &gRPC.CoordinatorMessage {
				coordID : voteID,
			}
			ack, _ := client.Elected(context.Background(), coordmsg)
		
		}
	
	empty := &gRPC.EmptyMessage{}
	return empty
}

func (s *Node) Coordinator(ctx context.Context, coordmsg gRPC.CoordinatorMessage) (*gRPC.EmptyMessage, error) {
	//If not voted then we've completed a lap and must start a new election, otherwise:
	//If winner, "go into critical section" and set id to 0. If not winner, increase id by 1.
	//Finally, set voted to false
	if !s.voted {
		msg := &gRPC.ElectionMessage {
			topcnID: id,
		}
		ack, _ := client.Election(context.Background(), msg)
	} else if (coordmsg.coordID == id) {
		//This node is the winner. The critical section is accessed.
		log.Printf("%s is entering critical section", myName)
		id=0
		log.Printf("%s has left critical section and is now in back of queue", myName)
			msg := &gRPC.CoordinatordMessage {
				coordID: coordmsg.id,
			}
			ack, _ := client.Coordinator(context.Background(), msg)
	} else {
		//This node is not the winner. Elected() passes a message along to the next node in the circle
		id++
		electedMessage := &gRPC.CoordinatorMessage {
			coordID : coordmsg.id,
		}
		ack, _ := client.Coordinator(context.Background(), electedMessage)
	}
		
	empty := &gRPC.EmptyMessage{}
	return empty
}

// sets the logger to use a log.txt file instead of the console
func setLog() *os.File {
	// Clears the log.txt file when a new server is started
	if err := os.Truncate("serverlog.txt", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	// This connects to the log file/changes the output of the log informaiton to the log.txt file.
	f, err := os.OpenFile("serverlog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return f
}