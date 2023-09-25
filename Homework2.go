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

	fmt.Printf("Client sends syn...\n")
	c <- packet{x:clientSeq,y:0,}
	ack:= <- c
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
		c <- packet{x:1,y:msgLen, data:msg[1]}
		c <- packet{x:0,y:msgLen, data:msg[0]}
	}
}


func server (c chan packet) {
	//3-way handshake
	var clientSeqP packet
	var serverSeq int = rand.Intn(100)+1
	var clientAckP packet
	//data receival
	var msgLength int
	var bArray []byte
	var fullmsg string

	//3-way handshake
	clientSeqP = <- c
	fmt.Printf("Server recieved seq from client...\n")
	c <- packet{clientSeqP.x+1, serverSeq, 0}
	clientAckP = <- c
	fmt.Printf("Server recieved ack from client...\n")
	if (clientAckP.y == serverSeq + 1) {
		fmt.Printf("seqY has been recieved and sent back to server succesfully\n")
	}

	//data receival
	for {
		part := <- c
		msgLength = part.y

		if (len(bArray) != msgLength) {
			bArray = make([]byte, msgLength)
		}

		bArray[part.x] = part.data

	}

	//ordered construction of string
		if (len(bArray) == msgLength) {
			fullmsg = string(bArray)
			fmt.Printf("Full message: %s\n", fullmsg)
		}
	
}

func main() {
	comm:= make(chan packet)
	go client(comm)
	go server(comm)
	time.Sleep(4 * time.Second)
}
