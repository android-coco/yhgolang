package c

/*
#include "util.h"
*/
import "C"

import "fmt"

func GoSum(a, b int) {
	s := C.sum(C.int(a), C.int(b))
	fmt.Println(s)
}
