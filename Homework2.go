package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
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
		fmt.Printf("Client sends back with seqY+1...!\n")
		c <- packet{x:ack.x,y:ack.y+1}
		
		var msg string = "Hi"
		var msgLen int = len(msg)
		//using x field as sequence num
		c <- packet{x:0,y:msgLen, data:msg[0]}
		c <- packet{x:1,y:msgLen, data:msg[1]}
	} else {
		fmt.Printf("Received "+ strconv.Itoa(ack.x)+ ", expected "+ strconv.Itoa(clientSeq+1))
	}
}


func server (c chan packet) {
	var clientSeqP packet
	var serverSeq int = rand.Intn(100)+1
	var clientAckP packet
		
	clientSeqP = <- c
	fmt.Printf("Server sends x+1 to client...\n")
	c <- packet{clientSeqP.x+1, serverSeq, 0}
	fmt.Printf("Server received ack from client...\n")
	
	select {
	case clientAckP = <- c:
		case <- time.After(1 * time.Second):
			fmt.Printf("Too slow! Lost message! Listening...\n")
			server(c)
			return
	}
	

	if (clientAckP.y == serverSeq + 1) {
		fmt.Printf("seqY has been received and sent back to server succesfully\n")
	} else {fmt.Printf("Received "+ strconv.Itoa(clientAckP.y)+", expected "+ strconv.Itoa(serverSeq+1))}


	
}

func middleware(client chan packet, server chan packet) {
	var handshakedone bool = false
	for {
		var random int = rand.Intn(6)
		select {	
			case clientmsg:=  <- client: {
					if (random!=1 || handshakedone) {
					server <- clientmsg
					}
					if (random==1) {
						fmt.Printf("1  NOOOOOO\n")
						//time.Sleep(2 * time.Second)
					}
				}

			case servermsg := <- server: {
					if (random!=1 || handshakedone) {
						client <- servermsg
						handshakedone = true
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
