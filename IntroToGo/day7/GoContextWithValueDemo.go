package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
- can be used :
  - to pass a value to the function
  - cannot cancel a function
*/
func main() {
	//create a context to manage the long process f1
	//since main is the caller of f1 ,create it inside main()
	//step1: create the base context
	var contextVar = context.Background()
	//step2: create a derived context from the base context
	//method2: to generate a context
	context1 := context.WithValue(contextVar, "userId", 2255)

	var wg sync.WaitGroup
	wg.Add(2)
	//pass context1 to goroutines as an argument preferable first argument
	go f1(context1, &wg)
	go f2(context1, &wg)
	time.Sleep(4 * time.Second)
	wg.Wait()
}
func f1(contextArg context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-contextArg.Done():
		fmt.Println("f1 is cancelled")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Inside f1. Value passed is:", contextArg.Value("userId"))
		return
	}

	// time.Sleep(10 * time.Second) //delay in f1 and takes long time to complete. therefore can be closed using CONTEXT
	// fmt.Println("Inside f1")
	// wg.Done()
}
func f2(contextArg context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-contextArg.Done():
		fmt.Println("f2 is cancelled")
		return
	case <-time.After(3 * time.Second):
		fmt.Println("Inside f2. Value passed is:", contextArg.Value("userId"))
		return
	}
	// fmt.Println("Inside f2")
	// wg.Done()

}
