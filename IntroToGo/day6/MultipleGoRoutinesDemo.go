package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to multiple go routine")
	ch1 := make(chan string)
	//sending data
	for i := 1; i < 4; i++ {
		go func(id int) {
			ch1 <- fmt.Sprintf("Hello from goroutine %d ", id)
		}(i)
	}
	//receiving data
	for i := 1; i < 4; i++ {
		msg := <-ch1
		fmt.Println(msg)
	}

	time.Sleep(2 * time.Millisecond)

}
