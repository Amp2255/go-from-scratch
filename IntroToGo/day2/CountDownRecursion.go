package main

import "fmt"

func main() {
	fmt.Println("Count down starts now !")
	countDown(10)
}

func countDown(i int) {
	if i > 0 {
		fmt.Println(i)
		i--
		countDown(i)
	} else {
		fmt.Println(0)
		fmt.Println("D day!")
	}
}
