package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// waitgroup to maintain order or synchronize goroutines///wg()
func main() {
	var wg sync.WaitGroup
	wg.Add(2) //specifies the number of goroutines the main() & wg has to consider
	fmt.Println("Initial statement from MAIN")
	go printLetters(&wg)
	go printNumbers(&wg)
	fmt.Println("Number of goroutines running at the moment :", runtime.NumGoroutine())
	wg.Wait() //waits for done() to be called inside goroutines
	fmt.Println("Final statement from MAIN")

}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() //using defer, the done() will be called at the end of execution of this function
	for i := 1; i <= 10; i++ {
		fmt.Println("Number :", i)
		time.Sleep(5 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done() //using defer, the done() will be called at the end of execution of this function
	for i := 'A'; i <= 'G'; i++ {
		fmt.Printf("Letter : %c\n", i)
		time.Sleep(5 * time.Millisecond)
	}
}
