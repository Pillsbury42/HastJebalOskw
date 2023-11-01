package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
)

func fork (i int, f chan bool) {
	for {
		f<-true
		<- f
	}
}

func philo (i int, l chan bool, r chan bool) {
	leftHand := false
	rightHand := false
	ate := 0

	for ate < 5{
		select {
			case leftHand = <- l:
			case rightHand = <- r:
		}
		if (leftHand && rightHand) {
			r<-true
			l<-true
			fmt.Printf("%d done eating. Has  eaten %d times \n", i, ate+1)
			ate++
			leftHand = false
			rightHand = false
			//reset order the forks are grabbed in
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	
	free := make([]chan bool, 5)

	for i := 0; i < 5; i++ {
		free[i] = make(chan bool)
	}
	
	for i := 0; i < 4; i++ {
		go fork(i, free[i])
		go philo(i, free[i], free[i+1])
	}

	//last fork
	go fork(4, free[4])
	//this one grabs forks in the other direction than the rest
	go philo(4, free[0], free[4])


	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}