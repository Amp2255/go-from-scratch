package main

import (
	"fmt"
	"sync"
)

func main() {
	//nested func sample
	greet("CDAC")

	//function returning fun
	res := result("Hello World !") //assign the fn to a var
	res()                          //invoke the func by calling the var()

	//CLOSURES IN GO
	sample := sampleClosure()
	fmt.Println(sample())

}

// nested functions
func greet(name string, wg *sync.WaitGroup) { //outer function
	var showName = func() { //inner function -- inner function SHOULD BE an anonymous function
		fmt.Println("Name is :", name)
		fmt.Println("Welcome ", name) //name which is an argument of outer function
		// can be accessed by inner function
	}
	showName()
}

// function returning another function
func result(s string) func() {
	return func() {
		fmt.Println(s, "This is a function returned by another function!")
	}
}

// CLOSURES IN GO
func sampleClosure() func() string {
	name := "CDAC"
	//this will return an anonymous func
	return func() string {
		name = " Hi ..." + name
		return name
	}
}
