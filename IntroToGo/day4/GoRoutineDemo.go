package main

import (
	"fmt"
	"sync"
	"time"
)

//NOTE: check parallelism VS concurrency

/*
Thread is a light weight process and
GO Routine is a very lightweight thread
	- light weight
	- independently executing function
	- managed by go runtime
	- cheaper than normal thread
	- use sync for maintatining the order else scheduler will select the thread in random
		- wight group wg

*/

// keyword 'go' to mark it as a go routine
// it will be executed only at the end
func main() {
	var wg sync.WaitGroup

	go greet("Hello ", &wg)
	//greet("World!", &wg)
	go greet("This is confusing", &wg)
	go greeterInfinite("Im an infinite", &wg)

	greetWithWeightGroup("I hav wt group", &wg)
	wg.Wait()
	fmt.Println("Final statement ")

}

func greet(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond) // this sleep will be applied to the function if it has prefix 'go' while invokation
		fmt.Println(s)
	}
}

func greeterInfinite(s string, wg *sync.WaitGroup) {
	for {
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

func greetWithWeightGroup(s string, wg *sync.WaitGroup) {
	for {
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}
