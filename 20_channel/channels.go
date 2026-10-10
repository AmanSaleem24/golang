package main

import (
	"fmt"
	"time"
)

func processChan(numChan chan int){
	for num := range numChan {
		fmt.Println("Processing the numChan with the value of", num)
		time.Sleep(time.Second)
	}
}

func emailSender(emailChan <-chan string, done chan<- bool){
	defer func() { done <- true} ()


	for email:= range emailChan {
		fmt.Println("Sending mail to", email)
		time.Sleep(time.Second*2)
	}

}

func main(){

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	} ()

	go func() {
		chan2 <- "golang"
	} ()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val :=  <- chan1:
			fmt.Println("Chan1 value is", chan1Val)
		case chan2Val := <- chan2:
			fmt.Println("Chan 2 value is", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)


	// for i:= range 100 {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// close(emailChan)

	// <- done


	// numChan := make(chan int)

	// go processChan(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "ping"

	// msg := <- messageChan

	// fmt.Println(msg)
}
