package main

import (
	"fmt"
	"slices"
)

// SLICE :: undefined sized array => var slice_one = []int{1,2,3,4}
func main() {

	var slice_one = []int{1, 2, 3, 4, 5}
	fmt.Println("The slice elements are :", slice_one)
	fmt.Println("The slice size :", len(slice_one))

	//slice from an array
	var arr = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("The array elements are :", arr)
	slice_two := arr[0:5]
	fmt.Println("The slice elements are slice_two:", slice_two)

	slice_three := arr[1:5]
	fmt.Println("The slice elements are slice_three:", slice_three)

	//make() to create a slice
	//numbers := make([]int,5,7)

	//append() add elements to a slice
	slice_two = append(slice_two, slice_three...)
	fmt.Println("The slice elements are slice_two:", slice_two)

	//copy() copy one slice to another
	source := []int{2, 3, 5, 7}
	destination := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("number of elements copied to destination :", copy(destination, source))
	fmt.Println("destination :", destination)

	//equal() to compare two slices
	fmt.Println("destination and source are equal : ", slices.Equal(destination, source))

}
