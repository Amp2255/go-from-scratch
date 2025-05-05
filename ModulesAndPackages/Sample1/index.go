package main

import (
	"SampleProject/package1"
	printPackage "fmt" //package name alias. can be used when the package name is very long
)

func main() {
	printPackage.Println("The variable accessed from package1 is :", package1.RevString)
}
