package main 

import (
	"fmt"
)

func main() {
	var i int = 15;
    fmt.Printf("%d 的阶乘是 %d\n", i, recursion1(i));

    for i := 0; i < 10; i++ {
       fmt.Printf("%d\t", fibonacci(i));
    }
    fmt.Printf("\n");
}

func recursion(x int) (y int) {
	relust := 1;
	if x == 0 {
		relust = 1;
	}else {
		relust = recursion(x - 1) * x;
	}
   return relust;
}

func recursion1(x int) (y int) {
	if x == 0 {
		y = 1;
	}else {
		y = recursion(x - 1) * x;
	}
   return;
}

func fibonacci(n int) (x int) {
  if n < 2 {
   return n;
  }
  return fibonacci(n-2) + fibonacci(n-1);
}