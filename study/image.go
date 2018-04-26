package main

/**
练习：图像
还记得之前编写的图片生成器吗？我们再来编写另外一个，不过这次它将会返回一个 image.Image 的实现而非一个数据切片。

定义你自己的 Image 类型，实现必要的方法并调用 pic.ShowImage。

Bounds 应当返回一个 image.Rectangle ，例如 image.Rect(0, 0, w, h)。

ColorModel 应当返回 color.RGBAModel。

At 应当返回一个颜色。上一个图片生成器的值 v 对应于此次的 color.RGBA{v, v, 255, 255}。
 */
import (
	"image"
	"fmt"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct {
	W int
	H int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) ColorModel() color.Model {
	return color.RGBA64Model
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(100, 100).RGBA())

	m1 := Image{200, 200}
	pic.ShowImage(m1)
}
