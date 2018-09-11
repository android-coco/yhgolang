package main

import (
	"fmt"
	"os"
	"bufio"
	"awesomeProject/mkw/functional/fib"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			break
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE, 0666)
	//自定义err
	//err = errors.New("this is a custom error")
	if err != nil {
		if pathError,ok := err.(*os.PathError); !ok{
			panic(err)
		}else{
			fmt.Printf("%s, %s , %s \n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		//fmt.Println("file already exists")
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	fib := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, fib())
	}
}
func main() {
	//tryDefer()
	writeFile("fib.txt")
}
