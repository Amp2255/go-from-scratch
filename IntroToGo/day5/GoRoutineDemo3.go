package main

import (
	"fmt"
	"runtime"
	"time"
)

// sleep() will make the main method to wait of go sleep so that we could get the go routines completed by that time
// else main () will get executed and will close and we wont be able to make sure whether the goroutines were run successfully or not
// sleep() prevents early exit() of main()
func main() {
	fmt.Println("Initial statement from MAIN")
	go printLetters()
	go printNumbers()
	normalFunc()
	fmt.Println("Number of goroutines running at the moment :", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	fmt.Println("Final statement from MAIN")

}

func normalFunc() {
	fmt.Println("Hi, I am a normal function!")
}

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Number :", i)
		time.Sleep(5 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'A'; i <= 'G'; i++ {
		fmt.Printf("Letter : %c\n", i)
		time.Sleep(5 * time.Millisecond)
	}
}
