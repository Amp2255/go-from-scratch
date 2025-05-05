package main

import "fmt"

func main() {
	sum(1, 12, 3, 5, 6, 6, 8)
}

func sum(numbers ...int) {
	for i, num := range numbers {
		fmt.Println(" numbers :", i, num)

	}
}
