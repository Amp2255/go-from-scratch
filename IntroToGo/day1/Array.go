package main

import "fmt"

func main() {
	//array of numbers with size 5
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("The array elements are :", arr)
	fmt.Println("The array size :", len(arr))

	//array of string with undefined size
	var strArr = [...]string{"anju", "jonathan", "anakha", "suraja"}
	fmt.Println("The array elements are :", strArr)
	fmt.Println("The array size :", len(strArr))

	//array of string with shorthand notation
	shorthandArray := [...]string{"anju", "jonathan", "anakha", "suraja"}
	fmt.Println("The array elements are :", shorthandArray)
	fmt.Println("The array size :", len(shorthandArray))

	//initialize specific elements of an array
	specific_Elements_Array := [5]string{0: "anju", 3: "jonathan", 4: "suraja"}
	specific_Elements_Array[2] = "anakha"
	fmt.Println("The array elements are :", specific_Elements_Array)
	fmt.Println("The array size :", len(specific_Elements_Array))

	//loop through an array
	var array_Loop = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array_Loop); i++ {
		fmt.Println(array_Loop[i])
	}

	//sum of an array
	array_Loop = [5]int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i < len(array_Loop); i++ {

		sum = sum + array_Loop[i]
	}
	fmt.Println(sum)

}
