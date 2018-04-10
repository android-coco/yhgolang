// Example code for Chapter ? from "Build Web Application with Golang"
// Purpose: Hello world example demonstrating UTF-8 support.
// To run in the console, type `go run main.go`
// You're missing language fonts, if you're seeing squares or question marks.
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちは世界\n")
    HOME:= os.Getenv("GOPATH")
    fmt.Println(HOME)
    fmt.Printf(os.Getenv("JAVA_HOME"))
    fmt.Println(os.Getenv("USER"))
    fmt.Println("no value for $USER")
    fmt.Println(throwsPanic(A))

    var str1 str
    str1.name,str1.age = "1",12
    a := str{"2",15,human{"2",15,20}}
    fmt.Println(a,str1)
    fmt.Println(a.human.name,a.name,a.int)
}
type human struct {
    name string
    age int
    int
}
type str struct {
    name string
    age int
    human
}

func throwsPanic(f func()) (b bool) {
    defer func() {
        if x := recover(); x != nil {
            b = true
        }
    }()
    f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
    return
}
func A()  {
    panic("chu cuo")
}
