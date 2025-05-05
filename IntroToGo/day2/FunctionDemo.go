package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter the numbers")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	sum(a, b)
	res := result(a, b)
	fmt.Println("Sum from the function res :", res)

	fmt.Println("Returned values from the multipleValues() :")
	i, s := multipleValues(a, b)
	fmt.Println(" i is ", i, " and s is ", s)
}

// call a function
func sum(a int, b int) {
	fmt.Println("Sum of the numbers :", a+b)
}

// function returning a value
func result(a, b int) int {
	return (a + b)
}

// function returning multiple values
func multipleValues(a, b int) (int, string) {
	total := a + b
	cal := "sum"
	return total, cal
}
