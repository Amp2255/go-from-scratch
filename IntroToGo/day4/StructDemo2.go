package main

import "fmt"

// nested structures
func main() {

	fmt.Println("Hello Go Nested Struct!")

	type Address struct {
		city    string
		pincode int
	}
	type Person struct {
		name     string
		age      int
		diabetic bool
		address  Address //nested struct
	}

	var person1 = Person{
		name:     "Lords",
		age:      33,
		diabetic: true,
		address: Address{
			city:    "Lille",
			pincode: 59800,
		},
	}
	fmt.Println("Person 1 details :", person1)
	fmt.Println("Person 1 city :", person1.address.city)
}
