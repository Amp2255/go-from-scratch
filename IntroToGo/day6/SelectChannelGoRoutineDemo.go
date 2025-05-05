package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Select Channel")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Millisecond)
		ch1 <- "Data from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Millisecond)
		ch2 <- "Data from ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timed Out")
	case msg := <-ch2:
		fmt.Println(msg)
	}
}
