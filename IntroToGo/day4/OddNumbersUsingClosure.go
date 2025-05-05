package main

import "fmt"

func main() {
	//print odd numbers from 0 to n
	n := 110
	oddNumbers(n)
}

func oddNumbers(n int) {

	var oddNums = func() {
		for i := 0; i < n; i++ {
			if (i % 2) != 0 {
				fmt.Println(i)
			}
		}
	}
	oddNums()
}
