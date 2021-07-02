package main

import (
	"encoding/json"
	"fmt"
)

type JsonTest struct {
	X string      `json:"x"`
	Y json.Number `json:"y"`
}

type ConfigOne struct {
	Daemon string
}
func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}


func main() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	println(a, b, c,x)
	return
	switch 1 {
	case 1, 2, 3:
		fmt.Println("111111")
	}

	for i, j := 0, 0; i <= 5 && j <= 5; i, j = i+1, j-1 {
	}

	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)

	s := `{"x":"golang", "y":520.0}`
	var jt JsonTest

	err := json.Unmarshal([]byte(s), &jt)
	if err == nil {
		fmt.Printf("%+v\n", jt)
	} else {
		fmt.Println(err)
		fmt.Printf("%+v\n", jt)
	}
}
