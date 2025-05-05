package main

import "fmt"

func main() {
	fmt.Println(" Anonymous function sample")

	//anonymous function
	var greet = func() {
		fmt.Println("Hello I am anonymous")
	}
	greet()

	//pass arguments to anonymous function
	var greet1 = func(str string) {
		fmt.Println("Hello " + str + " I am anonymous")
	}
	greet1("world")

	//returning values from an anonymous function
	var sum = func(a, b int) int {
		total := a + b
		return total

	}
	result := sum(2, 3)
	fmt.Println("Sum is :", result)

	//anonymous function as arguments to another function
	addition := func(n1, n2 int) int {
		return n1 + n2
	}
	squareVal := findSquare(addition(1, 4))
	sumOfNumbers := addition(1, 4)
	fmt.Println("Square of :", sumOfNumbers, " is :", squareVal)
}

func findSquare(num int) int {
	square := num * num
	return square
}
