package main

import "fmt"

func main() {
	//array of numbers with size 5
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("The array elements are :", arr)
	fmt.Println("The array size :", len(arr))