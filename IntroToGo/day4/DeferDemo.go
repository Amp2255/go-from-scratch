package main

import "fmt"

// defer keyword to delay the execution of a process
// delay by waiting for main to return ie,
// it will be eecuted only at the end
// if there are multiple defers it takes the LIFO order
func main() {
	fmt.Println("Defer demo")
	fmt.Println("Start")
	defer fmt.Println("Stmt 1")
	fmt.Println("Stmt 2")
	defer fmt.Println("End")
}
