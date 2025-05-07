package package1

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

var RevString string = "PrintMeIfUCan"

func ReverseString() string {
	str := " A sample string to reverse"

	rev := reverse.String(str)
	RevString = rev

	fmt.Println("The reverse of \n", str, "\n is\n", rev)
	return rev
}
