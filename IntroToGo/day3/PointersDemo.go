package main

import "fmt"

func main() {

	var roll int = 10
	fmt.Println(&roll, "   ", roll)

	//declare a pointer variable
	var ptr *int
	fmt.Println("before initialization,ie, default value of pointer  :", ptr)
	ptr = &roll
	fmt.Println("after initialization  :", ptr)
	fmt.Println("value held in the address stored in ptr :", *ptr)
	fmt.Println("address of ptr :", &ptr)

	//pass pointer to a function
	var num = 100
	fmt.Printf("value before update() call :", num)
	update(&num)
	fmt.Printf("value after update() call :", num)

	//function returning a pointer
	str := giveAPointer()
	fmt.Println(" value returned by giveAPointer() using pointer :", *str)

	//take and return pointer
	fmt.Println("value of num before func :", num)
	result := recieveAndGivePointer(&num)
	fmt.Println("returned value from func :", result)
	fmt.Println("content of returned value from func :", *result)

	//call by value & call by reference
	callByVar := 1
	fmt.Println(" value of callByVar in main()  ", callByVar)
	callByValue(callByVar)
	fmt.Println(" value of callByVar in main() after callByValue  ", callByVar)
	callByRef(&callByVar)
	fmt.Println(" value of callByVar in main() after callBYRef ", callByVar)

}

func update(num *int) {
	*num = 300
}

func giveAPointer() *string {

	res := "hello"
	return &res
}

func recieveAndGivePointer(number *int) *int {

	*number = 1000
	return number
}

func callByValue(callByVar int) {
	callByVar = 11
	fmt.Println("value in call by val ", callByVar)
}

func callByRef(callByVar *int) {
	*callByVar = 13
	fmt.Println("value in call by ref ", *callByVar)
}
