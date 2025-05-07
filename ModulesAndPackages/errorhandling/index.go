package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := f1(233)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}

}

func f1(i int) (int, error) {
	if i < 10 {
		return 100, nil
	} else {
		//user defined error
		//where errors is an interface
		// var err error = errors.New("Too big value for i")
		//error description should be in capital case
		err := errors.New("TOO BIG VALUE FOR i")
		return -1, err
	}
}
