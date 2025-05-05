package main

import (
	"fmt"
)

func main() {
	var arr = [5][2]int{{0, 1}, {2, 8}, {3, 6}, {4, 6}, {1, 2}}
	fmt.Println("The array elements are :", arr)

	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr[i][j])
		}
	}
	fmt.Println("The array size :", len(arr))

	//RANGE ::: for iterating over an array

	for index, item := range arr {
		fmt.Printf("number [%d]=[%d]\n", index, item)
		fmt.Println(index, " has ", item)
	}

	//RANGE ::: for iterating over an array/string

	for k, ch := range "welcome" {
		fmt.Printf("%#U starts at byte position %d\n", ch, k)
	}
}
