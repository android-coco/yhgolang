package main

import (
	"fmt"
)

type TZ int

func (tz *TZ) print(a int) {
	*tz += TZ(a)
}

func main() {
	var tz TZ
	tz.print(100)
	fmt.Println(tz)
}
