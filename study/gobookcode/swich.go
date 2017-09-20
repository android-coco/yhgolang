package main

import (
	"fmt"
)

func main() {
	var i = 0
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fallthrough
	case 4:
		fmt.Println(4)
	case 5,6,7:
		fmt.Println(5,6,7)
	default:
		fmt.Fprintln("Defult")
	}
}

