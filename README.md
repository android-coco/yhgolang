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


