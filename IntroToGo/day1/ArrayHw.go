package main

import "fmt"

// output all pairs of elements from array 'arr' of size n whose sum is equal to 'b'

func main() {
	var n, b int
	fmt.Println("Enter the size of array")
	fmt.Scanln(&n)
	fmt.Println("Enter the sum to check")
	fmt.Scanln(&b)
	fmt.Println("Enter the elements of array")
	arr := [5]int{}
	i := 0
	for i < 5 {
		fmt.Scanln(&arr[i])
		i++
	}
	fmt.Println("Array elements are :", arr)
	i = 0
	for i < 5 {
		j := i + 1
		for j < 5 {
			if (arr[i] + arr[j]) == b {
				fmt.Println(arr[i], ",", arr[j])
			}
			j++
		}
		i++
	}
}
