package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operstion: %s", op)
	}
}

func applay(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d)\n", opName, a, b)
	return op(a, b)
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for _, v := range numbers {
		s += v
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("Error handling")
	if relust, err := eval(3, 4, "d"); err != nil {
		fmt.Println("Error:",err)
	} else {
		fmt.Println(relust)
	}

	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n",q, r)

	fmt.Println(applay(pow, 3, 4))
	fmt.Println("pow(3, 4) is:",applay(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))

	}, 3, 4))

	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a,b = swap(a, b)
	fmt.Println("a, b after swap is:", a, b)
}
