package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	//"strconv"
	"strings"

	// this has to be the same as the go.mod module,
	// followed by the path to the folder the proto file is in.
	gRPC "github.com/hannaStokes/handin5/gRPC"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientsName = flag.String("name", "default", "Senders name")
var serverPorts = [4]string{"5400", "5401", "5402", "5403"}

var server gRPC.AuctionClient //the server
var ServerConn *grpc.ClientConn  //the server connection

func main() {
	//parse flag/arguments
	flag.Parse()

	fmt.Println("--- CLIENT APP ---")

	//log to file instead of console
	f := setLog()
	defer f.Close()

	//connect to server and close the connection when program closes
	fmt.Println("--- join Server ---")
	ConnectToServer()
	defer ServerConn.Close()

	msg := &gRPC.BidMessage{
		ClientName: *clientsName,
		Timestamp:  clientsTime,
	}

	go bid(ctx, msg)

	//start the bidding
	parseInput()
}

func (s *Server) bid(ctx context.Context, msg *gRPC.BidMessage) (*gRPC.BidReplyMessage, error)

func ConnectToServer() {

	//dial options
	//the server is not using TLS, so we use insecure credentials
	//(should be fine for local testing but not in the real world)
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	//dial the server, with the flag "server", to get a connection to it
	log.Printf("client %s: Attempts to dial on port %s\n", *clientsName, *serverPort)
	conn, err := grpc.Dial(fmt.Sprintf(":%s", *serverPort), opts...)
	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return
	}
	//fmt.Printf("Error here")
	// makes a client from the server connection and saves the connection
	// and prints rather or not the connection was is READY
	server = gRPC.NewAuctionClient(conn)
	ServerConn = conn
	log.Println("the connection is: ", conn.GetState().String())
}

func parseInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type the message you wish to send below")
	fmt.Print("-> ")


	//Infinite loop to listen for clients input.
	for {

		//Read input into var input and any errors into err
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
	//input = strings.TrimSpace(input) //Trim input

		if !conReady(server) {
			log.Printf("Client %s: something was wrong with the connection to the server :(", *clientsName)
			continue
		}

		//terminal input parsing here
		matchresult, err := regexp.MatchString("bid (\\d+)", input)
		if (input == "result"){
			fmt.Printf("Getting result...")
			result()
		} else if (matchresult){
			fmt.Printf("Processing bid...")
			var splitinput = strings.Split(input, " ")
			inputint, _ := strconv.Atoi(splitinput[1])
			bid(inputint)
			//print depending on bidreplymessage, ie. did it go through or was it lower than the current bid
		} else {
			fmt.Println("Unknown command. Type 'result' for current auction status. Type 'bid (integer)' to bid.")
		}

		if err != nil {
			log.Printf("Client %s: no response from the server, attempting to reconnect", *clientsName)
			log.Println(err)
		}
	}
}
func bid(bidamount int){

}

func result(){

}

// Function which returns a true boolean if the connection to the server is ready, and false if it's not.
func conReady(s gRPC.AuctionClient) bool {
	return ServerConn.GetState().String() == "READY"
}

// sets the logger to use a log.txt file instead of the console
func setLog() *os.File {
	f, err := os.OpenFile("log_"+*clientsName+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return f
}