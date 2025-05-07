package main

import "fmt"

func main() {
	f1()
	fmt.Println("main ending")
}

func f1() {
	fmt.Println("Hello World")
	fmt.Println("I am function f1")
	panic("panic demo")
	fmt.Println("After panic()")
}
