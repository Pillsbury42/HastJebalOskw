// Server/Node
package main

import (
	"fmt"

	// "io"
	"log"
	"net"

	// this has to be the same as the go.mod module,
	// followed by the path to the folder the proto file is in.
	gRPC "github.com/Pillsbury42/HastJebalOskw/handin5/gRPC"

	"google.golang.org/grpc"
)

func main() {
	//f := setLog()
	//defer f.Close()

	// This parses the flags and sets the correct/given corresponding values.
	fmt.Println(".:Node is starting:.")

	// launch the server
	go launchServer(port)

	// code here is unreachable because launchServer occupies the current thread.
}
func launchServer(port string) {
	log.Printf("Attempting to create listener on port %s\n", port)

	// Create listener tcp on given port or default port 5400
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Printf("Failed to listen on port %s: %v", port, err) //If it fails to listen on the port, run launchServer method again with the next value/port in ports array
		return
	}

	// makes gRPC server using the options
	// you can add options here if you want or remove the options part entirely
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// makes a new server instance using the name and port from the flags.
	server := gRPC.UnimplementedAuctionServer

	gRPC.RegisterAuctionServer(grpcServer, server) //Registers the server to the gRPC server.

	log.Printf("Listening at %v\n\n", list.Addr())

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	// code here is unreachable because grpcServer.Serve occupies the current thread.
}
