package main

import (
	"fmt"
)

func main() {

	fmt.Println("Channels are bidirectional: it can send and receive data")
	ch1 := make(chan string)

	go func() {
		ch1 <- "Data from ch1" // send data to channel
	}()
	msg := <-ch1 // receive data from channel
	fmt.Println("message received from ch1 is:", msg)
}
