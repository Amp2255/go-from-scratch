package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to buffered channel!!")
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for msg := range ch {
		fmt.Println(" message ", msg)
	}

}
