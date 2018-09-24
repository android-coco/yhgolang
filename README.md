## 闭包-defer-painc-recover()
##  闭包

```go
//返回值为函数  闭包
func closure(x int) func(int) int {
    fmt.Printf("%p\n", &x)
    return func(y int) int {
        fmt.Printf("%p\n", &x)
        return x + y
    }
}
```

## painc  recover()

```go
func main() {
    J()
    K()
    L()
}
func J() {
    fmt.Println("Func J")
}
func K() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recover in K")
        }
    }()
    panic("Panic in K")//相当于抛出异常
}
func L() {
    fmt.Println("Func C")
}
```

## defer 用defer定义的函数 后定义的先执行 和栈一样 先进后出


```go
//分析运行结果
func main() {
    var fs = [4]func(){}
    for i := 0; i < 4; i++ {
        defer fmt.Println("defer i = ", i)                           //3,2,1,0
        defer func() { fmt.Println("defer_closure1 i=", i) }()       //4,4,4,4
        defer func(i int) { fmt.Println("defer_closure2 i=", i) }(i) //3,2,1,0
        fs[i] = func() {
            fmt.Println("closure i = ", i) //4,4,4,4
        }
    }

    for _, f := range fs {
        f()
    }
}
```

## _import下划线的作用
在Golang里，import的作用是导入其他package，但是今天在看beego框架时看到了import 下划线，不知其意，故百度而解之。

　　import 下划线（如：import _ hello/imp）的作用：当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import _ 引用该包。即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。
【实列】
 目录结构：
  GOPATH
      --bin
      --pkg
      --src
            --yhgolang
               --org.yh
                 --uitl
                     importDemo.go 
            --studyGo
              append.go
append.go
```go
package main

import (
    _ "yhgolang/org.yh/util"
    "fmt"
)

func main() {

    //util.Test()  编译报错，说：undefined: util in util.Test

    /**
            append主要用于给某个切片（slice）追加元素
            如果该切片存储空间（cap）足够，就直接追加，长度（len）变长；
            如果空间不足，就会重新开辟内存，并将之前的元素和新的元素一同拷贝进去
                第一个参数为切片，后面是该切片存储元素类型的可变参数
        基础用法：
    */
    slice := append([]int{1, 2, 3}, 4, 5, 6)
    fmt.Println(slice) //[1 2 3 4 5 6]
    //第二个参数也可以直接写另一个切片，将它里面所有元素拷贝追加到第一个切片后面。
    //要注意的是，这种用法函数的参数只能接收两个slice，并且末尾要加三个点
    slice1 := append([]int{1, 2, 3}, []int{4, 5, 6}...)
    fmt.Println(slice1) //[1 2 3 4 5 6]
    //还有种特殊用法，将字符串当作[]byte类型作为第二个参数传入
    //append函数返回值必须有变量接收，不然编译器会报错
    bytes := append([]byte("hello"), " world"...)
    fmt.Println(string(bytes))
}
输出结果：util-init() fadsfadsf.
```

```go
package util

import (
    "fmt"
)

func Test() int {
    return 1
}

func init() {
    fmt.Println("fadsfadsf")
}

```

