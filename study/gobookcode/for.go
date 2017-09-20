package main

import "fmt"

func main() {

	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				goto JLoop
			}
			fmt.Println(i)
		}
	}
JLoop:
}
