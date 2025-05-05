package main

import (
	"fmt"
	"sync"
)

// multiple senders but only one receiver
func main() {
	fmt.Println("Welcome to multiple go routine")
	ch1 := make(chan string)
	var wg sync.WaitGroup

	//3 go routines to send data
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ch1 <- fmt.Sprintf("Hello from goroutine %d ", id)
		}(i)
	}
	//close the channel after the send operations are over
	go func() {
		wg.Wait()
		close(ch1)
		fmt.Println("Channel closed ")
		ch1 <- fmt.Sprintf("Trying to send data to closed channel")
	}()
	//receiving data
	for msg := range ch1 {

		fmt.Println(msg)
	}

}
