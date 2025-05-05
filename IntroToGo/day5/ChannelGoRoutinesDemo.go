package main

import (
	"fmt"
)

// Channel
// both sending and receiving cannot be done using a single channel
// we need seperate channel for each operations
// once assigned , a specific channel can pass that type data only(if string then string only)
func main() {
	fmt.Println("How goroutines communicate??? \nThrough Channel!!")
	ch1 := make(chan string)
	// ch <- 10    //implies data 10 of type int is being send to channel ch
	// val := <-ch //implies ch sends data 10 to val
	go sendData(ch1)
	msg := <-ch1
	fmt.Println("\nValue recieved from channel is :", msg)
}

func sendData(sendCh chan string) {
	sendCh <- "Hello from goroutine"
}
