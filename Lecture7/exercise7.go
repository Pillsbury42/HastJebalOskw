package main

import (
	"fmt"
	"strconv"
	"time"
)

func setLeader(id int) {
	fmt.Println("Starting " + strconv.Itoa(id) + " at " + time.Now().String())
	time.Sleep(1 * time.Second)
	fmt.Println("Ending " + strconv.Itoa(id) + " at " + time.Now().String())
}
func node(id int, receiver chan bool, sender chan bool) {
	for {
		<-receiver
		setLeader(id)
		sender <- true
	}
}

func main() {

	done := make([]chan bool, 5)
	for i := 0; i < 5; i++ {
		done[i] = make(chan bool)
	}
	for i := 0; i < 4; i++ {
		go node(i, done[i], done[i+1])
	}
	go node(4, done[4], done[0])
	done[0] <- true
	for {

	}
	//fmt.Println("Finished")
}
