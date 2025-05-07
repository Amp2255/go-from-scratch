package main

import "fmt"

func main() {
	//uncomment f2 to check defer
	//f2() //defer will get executed after the panic()

	f3() //recover() demo
	fmt.Println("main ending")
}

func f2() {
	defer fmt.Println("f2 function completed") //defer will get executed after the panic()
	defer f4()
	fmt.Println("Hello World")
	fmt.Println("I am function f1")
	panic("panic demo")
	fmt.Println("After panic()")
}

func f3() {
	defer fmt.Println("f2 function completed") //defer will get executed after the panic()
	fmt.Println("Hello World")
	fmt.Println("I am function f1")
	panic("panic demo")
	fmt.Println("After panic()")
}

func f4() {
	message := recover()
	if message != nil {
		fmt.Println("*********", message)
	}
}
