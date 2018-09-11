package iota

import "fmt"

//跳值使用
const (
	a = iota
	b = iota
	_
	c = iota
	d = iota
)

//插队使用
const (
	a1 = iota
	b1 = 3.14
	c1 = iota
)

const (
	a2 = iota * 2
	b2 = iota
	c2 = iota
)

//隐式使用（当没有赋值时自动继承上面非空的规则）
const (
	a3 = iota * 2
	b3
	c3
)

//隐式使用（当没有赋值时自动继承上面非空的规则）
const (
	a4 = iota * 2
	b4 = iota * 4
	c4
	d4
)

//单行使用(c5沿用a5的规则，d5沿用b5的规则，而且未赋值的行元素要和之前的元素一致)
const (
	a5, b5 = iota + 1, iota + 3
	c5, d5
	f5     = iota
)

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println(a1, b1, c1)
	fmt.Println(a2, b2, c2)
	fmt.Println(a3, b3, c3)
	fmt.Println(a4, b4, c4, d4)
	fmt.Println(a5, b5, c5, d5, f5)
}
