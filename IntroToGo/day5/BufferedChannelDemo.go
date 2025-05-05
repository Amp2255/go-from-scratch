package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to buffered channel!!")
	ch := make(chan int, 2) //2 refers to number of channels to buffer
	ch <- 1
	ch <- 2
	fmt.Println("ch=1 :", <-ch)
	fmt.Println("ch=2 :", <-ch)
}
