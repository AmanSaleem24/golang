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

func emailSender(emailChan chan string, done chan bool){
	defer func() { done <- true} ()

	for email:= range emailChan {
		fmt.Println("Sending mail to", email)
		time.Sleep(time.Second*2)
	}

}

func main(){

	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)


	for i:= range 100 {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	<- done


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
