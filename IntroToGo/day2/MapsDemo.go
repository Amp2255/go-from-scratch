package main

import "fmt"

func main() {

	//create map with make()
	var makeMap = make(map[int]int)
	makeMap[1] = 100
	makeMap[2] = 200
	makeMap[3] = 300
	fmt.Println("print map elements :", makeMap)

	//create map without make()
	var mapLoop = map[int]int{1: 1, 2: 4, 3: 9}

	for k, v := range mapLoop {
		fmt.Printf("the square of %d is %d\n", k, v)
		fmt.Printf("----------\n")
	}

	//add elements
	mapLoop[5] = 25

	fmt.Println("map after addition", mapLoop)

	//remove elements
	delete(mapLoop, 5)
	fmt.Println("map after removing", mapLoop)

	//count occurrences of elements
	words := []string{"java", "golang", "react", "golang", "react"}
	wordCount := make(map[string]int)
	for _, v := range words {
		wordCount[v]++
	}
	fmt.Println("wordcount ", wordCount)
	for index, data := range words {
		wordCount[data]++
		fmt.Println("key : ", index, " value : ", data)
	}
	fmt.Println("wordcount ", wordCount)

	//maps used to cache results/ config details
	config := map[string]string{
		"host": "localhost",
		"port": "8080",
	}
	fmt.Println("config: -----------------", config)

	//maps to check for duplicates
	data := make(map[int]bool)
	nums := []int{1, 2, 3, 2, 5, 1, 2, 8, 6}
	for _, nums := range nums {
		if data[nums] {
			fmt.Println("Duplicate  :", nums)
		}
		data[nums] = true
		fmt.Println(data)

	}
}
