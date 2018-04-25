package main

/*
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
type error interface {
    Error() string
}
*/
import (
	"fmt"
	"time"
	"math"
)

func main() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
	fmt.Println(Sqrt(-2))
}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 	`error` 接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
				`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSprt float64

func (e ErrNegativeSprt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSprt(x)
	}
	if x == 0 || x == 1 {
		return x, nil
	}
	var z = float64(1)
	var oldz float64
	for {
		oldz = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(oldz, z)

		//if (oldz - z ) <= 0.00001 {
		//	break
		//}
		if math.Abs(oldz-z) < 1e-10 {
			break
		}
	}
	return z, nil
}
