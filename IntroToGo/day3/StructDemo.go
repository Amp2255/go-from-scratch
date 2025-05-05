package main

import "fmt"

func main() {
	fmt.Println("Hello Go Struct!")

	type Person struct {
		name     string
		age      int
		diabetic bool
	}

	person1 := Person{"John Doe", 43, false}
	fmt.Println("Details of person1 :", person1)

	var person2 = Person{
		name:     "Lords",
		age:      33,
		diabetic: true,
	}
	fmt.Println("Details of person2 :", person2)
	person3 := Person{"Peter Pan", 23, true}
	fmt.Println("Details of person3 :", person3)

	//accessing struct elements
	fmt.Println("----accessing struct elements-----")
	fmt.Println("Details of person1 :")
	fmt.Println("name :", person1.name,
		"\nage :", person1.age, "\ndiabetic :", person1.diabetic)
	fmt.Println("------------------")
	//array of struct
	fmt.Println("----array of struct-----")
	var people = [3]Person{}
	people[0] = person1
	people[1] = person2
	people[2] = person3
	fmt.Println("Details of People :", people)
	fmt.Println("------------------")

	//struct with a function as element
	fmt.Println("--------function as element of struct----------")
	type Citizen struct {
		name     string
		age      int
		isSenior func() bool
	}

	isSenior func(int) bool
	var citizen1 = Citizen{
		name: "Paul",
		age:  33,
		isSenior: func(int) bool {
			if 33 > 60 {
				return true
			} else {
				return false
			}
		}}

}
