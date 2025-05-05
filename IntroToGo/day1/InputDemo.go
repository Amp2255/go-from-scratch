package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter the name")
	fmt.Scanln(&name)
	fmt.Println("Name is :", name)
}
