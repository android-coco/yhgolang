package so

/*
#include "libzsign_so.h"
#cgo LDFLAGS: -ldl
*/
import "C"
import "fmt"

func Test() {
	i := C.CString("===")
	fmt.Println("20*30=", C.do_so_func(&i))
	var a,b = 1,2
	c,d := a,b
	fmt.Print(c,d)
}
