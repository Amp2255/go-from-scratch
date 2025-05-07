package main

import (
	test "SampleProject/package1"
	printPackage "fmt" //package name alias. can be used when the package name is very long
)

func main() {
	s := test.ReverseString()
	printPackage.Println("The function accessed from package1 :", s)
}
