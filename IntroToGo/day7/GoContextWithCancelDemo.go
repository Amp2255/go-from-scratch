package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
- can be used :
  - to cancel a go routine performing a very time consuming operation
  - to manage timeout across goroutines efficiently
  - for data sharing between go routines
*/
func main() {
	//create a context to manage the long process f1
	//since main is the caller of f1 ,create it inside main()
	//step1: create the base context
	var contextVar = context.Background()
	//step2: create a derived context from the base context
	//using which we can explicitly cancel long fns
	//method1: to generate a context along with a cancel function
	context1, cancelfn := context.WithCancel(contextVar)
	var wg sync.WaitGroup
	wg.Add(2)
	//pass context1 to goroutines as an argument preferable first argument
	go f1(context1, &wg)
	go f2(context1, &wg)
	time.Sleep(4 * time.Second)
	cancelfn() //it is function, when called will cancel all running goroutines
	wg.Wait()
}
func f1(contextArg context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-contextArg.Done():
		fmt.Println("f1 is cancelled")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Inside f1")
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
		fmt.Println("Inside f2")
		return
	}
	// fmt.Println("Inside f2")
	// wg.Done()

}
