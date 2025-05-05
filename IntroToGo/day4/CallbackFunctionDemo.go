package main

import "fmt"

func main() {
	var a int
	fmt.Println("Select 1 for addition, 2 for multiplication")
	fmt.Scanln(&a)
	if a == 1 {
		processNum(10, 25, add)
	} else {
		processNum(10, 25, multiply)
	}

}

func processNum(a int, b int, callback func(int, int) int) {
	res := callback(a, b)
	fmt.Println(res)
}

func add(n int, m int) int {
	return n + m
}

func multiply(n int, m int) int {
	return n * m
}
