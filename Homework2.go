package main

import (
	"fmt"
	"time"
	"math/rand"
)

type packet struct {
	x int
	y int
	data byte
}

func client (c chan packet) {
	var clientSeq int = rand.Intn(100)+1
	var ack packet

	fmt.Printf("Client sends syn...\n")
	c <- packet{x:clientSeq,y:0,}
	select {
		case ack = <- c:
		case <- time.After(3*time.Second):
			fmt.Printf("1  Too slow! Message lost. Retrying...\n")
			client(c)
		return
	}
	fmt.Printf("Client receives ack...\n")
	
	if ack.x==(clientSeq+1) {
		fmt.Printf("ack is correct!\n")
		fmt.Printf("Client sends back with seqY+1...\n")
		c <- packet{x:ack.x,y:ack.y+1}
		
		//data transmission
		//using x field as sequence num
		//and y as message length
		var msg string = "Hi"
		var msgLen int = len(msg)
		c <- packet{x:0,y:msgLen, data:msg[0]}
		c <- packet{x:1,y:msgLen, data:msg[1]}
	}
}


func server (c chan packet) {
	//3-way handshake
	var clientSeqP packet
	var serverSeq int = rand.Intn(100)+1
	var clientAckP packet
	//data receival
	// var msgLength int
	// var bArray []byte
	// var fullmsg string

	//3-way handshake
		
	//waiting to receive from client
	clientSeqP = <- c
	c <- packet{clientSeqP.x+1, serverSeq, 0}
	
	//waiting to receive from client again
	select {
		case clientAckP = <- c:
			fmt.Printf("Server recieved ack from client...\n")
		case <- time.After(1 * time.Second):
			fmt.Printf("Too slow! Lost message! Listening...\n")
			server(c)
			return
	}
	

	if (clientAckP.y == serverSeq + 1) {
		fmt.Printf("seqY has been received and sent back to server succesfully\n")
	}

	// //data receival
	// for {
	// 	part := <- c
	// 	msgLength = part.y

	// 	if (len(bArray) != msgLength) {
	// 		bArray = make([]byte, msgLength)
	// 	}

	// 	bArray[part.x] = part.data

	// }

	// //ordered construction of string
	// 	if (len(bArray) == msgLength) {
	// 		fullmsg = string(bArray)
	// 		fmt.Printf("Full message: %s\n", fullmsg)
	// 	}
	
}

func middleware(client chan packet, server chan packet) {
	

	for {
		var random int = rand.Intn(4)
		select {	
			case clientmsg:=  <- client: {
					if (random!=1) {
					server <- clientmsg
					}
					if (random==1) {
						fmt.Printf("1  NOOOOOO\n")
						//time.Sleep(2 * time.Second)
					}
				}

			case servermsg := <- server: {
					if (random!=1) {
					client <- servermsg
					}
					if (random==1) {
						fmt.Printf("2  NOOOOOOO\n")
						//time.Sleep(2 * time.Second)
					}
				}
		}
	}
}

func main() {
	clientcomm:= make(chan packet)
	servercomm:= make(chan packet)
	go client(clientcomm)
	go server(servercomm)
	go middleware(clientcomm, servercomm)
	time.Sleep(5 * time.Second)
}
