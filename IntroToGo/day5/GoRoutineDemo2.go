package main

import (
	"fmt"
	"time"
)

//NOTE: check parallelism VS concurrency

/*
Thread is a light weight process and
GO Routine is similar to thread scheduler in JAVA
	- light weight
	- independently executing function
	- managed by go runtime
	- cheaper than normal thread
	- use sync for maintatining the order else scheduler will select the thread in random
		- wight group wg

*/

// keyword 'go' to mark it as a go routine
// it will be executed only at the end
// we use wg. weight group to make main() to wait till the completion of execution of goroutines
func main() {
	fmt.Println("Initial statement from MAIN")
	go sayHello()
	go sayHello()
	go sayHello()
	go sayHello()
	go sayHello()
	fmt.Println("Final statement from MAIN")
	time.Sleep(time.Second)

}

func sayHello() {
	fmt.Println("Hello World ")
}
