package main

import "fmt"

func main() {
	arr_arg := [5]int{1, 2, 4, 5, 6}
	fmt.Println(" before update() call ", arr_arg)
	updateArray(&arr_arg)
	fmt.Println(" after update() call ", arr_arg)

	//assign pointer to pointer
	var x = 10
	var ptr = &x    //ptr to an int variable
	var pptr = &ptr //points to addr of a pointer ,ptr
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println("........................")
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(&ptr)
	fmt.Println("........................")
	fmt.Println(pptr)
	fmt.Println(&pptr)
	fmt.Println(*pptr)
	fmt.Println(**pptr)
	fmt.Println("........................")
}

func updateArray(arr_arg *[5]int) {
	for i := 0; i < len(arr_arg); i++ {
		arr_arg[i] *= 2
	}
}
