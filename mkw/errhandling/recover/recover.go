package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		err := recover()
		if err, ok := err.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			//panic(err)
			panic(fmt.Sprintf("I don't know what to do: %v", err))
		}
	}()
	//panic(errors.New("this is an error"))

	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	panic(123)
}

func main() {
	tryRecover()
}
