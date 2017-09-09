package main

import (
	"fmt"
)

// 常量
const (
	WHITH = iota
	BLACK
	BLUE
	RED
	YELLOW
)

// 自定义类型Color  其实是byte
type Color byte

// 自定义类型Box 结构体
type Box struct {
	width, height, depth float64
	color                Color
}

// 自定义类型切片元素为Box
type BoxList []Box //a slice of boxes

// 定义了接收者为Box，返回Box的容量。
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

/*
你也许会问SetColor函数里面应该这样定义*b.Color=c，
而不是b.Color=c，因为我们需要读取到指针相应的值。
你是对的，其实Go语言里面这两种方式都是正确的，
当你用指针去访问相应的字段时（虽然指针没有任何的字段），
Go语言知道你要通过指针去获取这个值，看到了吧，
Go语言的设计是不是越来越吸引你了。
*/
// 把Box的颜色改为c
func (b *Box) SetColor(c Color) {
	b.color = c
	//*b.color = c
}

// 定在在BoxList上面，返回list里面容量最大的颜色
func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITH) // 类型转换
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

/*
PaintItBlack里面调用SetColor的时候是不是应该写成(&bl[i]).SetColor(BLACK)，
因为SetColor的receiver是*Box，而不是Box。
你又对了，这两种方式都可以，
因为Go语言知道receiver是指针，
它自动帮你转了。也就是说，
如果一个method的receiver是*T，
你可以在一个T类型的实例变量V上面调用这个method，
而不需要&V去调用这个method。类似的，如果一个method的receiver是T，
你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method。
*/
// 把BoxList里面所有Box的颜色全部变成黑色。
func (bl BoxList) PaintltBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

// 定义在Color上面，返回Color的具体颜色（字符串格式）
func (c Color) String() string {
	strings := []string{"WHITH", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITH},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())
}
