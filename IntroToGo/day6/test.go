package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go f1(&wg)
	go f2(&wg)

	wg.Wait()
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("Inside f1")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	fmt.Println("Inside f2")
	wg.Done()
}
