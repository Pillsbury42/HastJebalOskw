package main

import (
	"bufio"

	//"container/list"
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
	gRPC "github.com/Pillsbury42/HastJebalOskw/handin5/gRPC"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientsName = flag.String("name", "default", "Senders name")
var serverPort = flag.String("serverp", "5400", "sending port")

var server gRPC.AuctionClient   //the server
var ServerConn *grpc.ClientConn //the server connection

var servermap map[int64]string

var leaderID int64

func main() {
	//parse flag/arguments
	flag.Parse()

	fmt.Println("--- CLIENT APP ---")

	//log to file instead of console
	//f := setLog()
	//Wdefer f.Close()

	servermap = make(map[int64]string)
	servermap[1] = "5400"
	servermap[2] = "5401"
	servermap[3] = "5402"

	leaderID = 1

	//connect to server and close the connection when program closes
	fmt.Println("--- join Server ---")
	ConnectToServer(*serverPort)
	defer ServerConn.Close()

	//start the bidding
	parseInput()
}

func ConnectToServer(port string) error {

	//dial options
	//the server is not using TLS, so we use insecure credentials
	//(should be fine for local testing but not in the real world)
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	//dial the server, with the flag "server", to get a connection to it
	fmt.Printf("client %s: Attempts to dial on port %s\n", *clientsName, port)
	conn, err := grpc.Dial(fmt.Sprintf(":%s", port), opts...)
	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return err
	}
	//fmt.Printf("Error here")
	// makes a client from the server connection and saves the connection
	// and prints rather or not the connection was is READY
	server = gRPC.NewAuctionClient(conn)
	ServerConn = conn
	fmt.Println("the connection is: ", conn.GetState().String())
	return nil
}

func parseInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type the command you wish to execute below")
	fmt.Print("Valid commands: bid <value>; result ")

	//Infinite loop to listen for clients input.
	for {

		//Read input into var input and any errors into err
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//input = strings.TrimSpace(input) //Trim input
		input = strings.Trim(input, "\r\n")

		if !conReady(server) {
			log.Printf("Client %s: something was wrong with the connection to the server :(", *clientsName)
			//continue
		}

		//terminal input parsing here
		matchresult, err := regexp.MatchString("bid (\\d+)", input)
		fmt.Printf(input)
		if input == "result" {
			fmt.Printf("Getting result...")
			result()
		} else if matchresult {
			fmt.Printf("Processing bid...")
			var splitinput = strings.Split(input, " ")
			inputint, _ := strconv.Atoi(splitinput[1])
			in := int64(inputint)
			bid(in)
			//print depending on bidreplymessage, ie. did it go through or was it lower than the current bid
		} else {
			fmt.Println("Unknown command. Type 'result' for current auction status. Type 'bid <integer>' to bid.")
		}

		if err != nil {
			log.Printf("Client %s: no response from the server, attempting to reconnect", *clientsName)
			log.Println(err)
		}
	}
}
func bid(bidamount int64) {

	msg := &gRPC.BidMessage{
		BidderName: *clientsName, //change proto
		Amount:     bidamount,
	}

	//not sure about this stuff with contextbackground. if it doesn't work, change ctx to context.Background() which is what it was before i messed with it
	res, err := server.Bid(context.Background(), msg)
	if err != nil || !conReady(server) {
		ServerConn.Close()
		//timeout code here, ie. send to random node
		for _, value := range servermap {
			err = ConnectToServer(value)
			if err == nil {
				break
			}
		}
		res, err = server.Bid(context.Background(), msg)

	}
	//If the leader is a different node from the one that answers, then also change server
	fmt.Println(res.LeaderID)
	fmt.Println(leaderID)
	if res.LeaderID != leaderID {
		leaderID = res.LeaderID
		fmt.Println("Wrong leader, closing connection")
		ServerConn.Close()

		for key, value := range servermap {
			fmt.Println("Trying another server")
			if key == leaderID {

				err = ConnectToServer(value)
				break
			}
		}
	}
	if res.Success == "Success" {
		fmt.Println("Bid completed. You are now the highest bidder.")
	} else if res.Success == "LowBid" {
		fmt.Println("Bid failed as it was too low.")
	} else if res.Success == "Over" {
		fmt.Println("Bid failed as the auction is over.")
	} else if res.Success == "Election" {
		fmt.Println("An error occured while bidding. Please try again.")
	}
}

func result() {

	res, err := server.Result(context.Background(), &gRPC.EmptyMessage{})
	if err != nil || !conReady(server) {
		fmt.Println("result error nil entered")
		//timeout code here, ie. send to random node
		ServerConn.Close()

		for _, value := range servermap {
			fmt.Println("Connecting to new server")
			err = ConnectToServer(value)
			if err == nil {
				break
			}
			fmt.Println("Exiting for loop in servermap")
		}
		res, err = server.Result(context.Background(), &gRPC.EmptyMessage{})
		fmt.Println("PRINTING ERROR HERE")
		fmt.Println(err.Error())
	}

	//If the leader is a different node from the one that answers, then also change server
	fmt.Println(leaderID)
	fmt.Println(res.LeaderID)
	if res.LeaderID != leaderID {
		leaderID = res.LeaderID
		ServerConn.Close()
		for key, value := range servermap {
			if key == leaderID {
				err = ConnectToServer(value)
				break
			}
		}
	}
	if res.Success == "Over" {
		fmt.Printf("Auction is over. The highest bid was %s by bidder %d", res.WinnerName, res.Highest)
	} else if res.Success == "NotOver" {
		fmt.Printf("Auction is NOT over. The highest bid so far is %s by bidder %d", res.WinnerName, res.Highest)
	} else if res.Success == "Election" {
		fmt.Println("An error occured while requesting results. Please try again.")
	}
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
