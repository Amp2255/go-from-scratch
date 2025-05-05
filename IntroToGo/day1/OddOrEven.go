package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number")
	fmt.Scanln(&num)
	if (num % 2) == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is odd")
	}
}
