package main
/* range 关键字用于for循环中迭代数组(arrays)、
	切片(slice)、通道(channel)或集合(map)的元素。
	在数组和切片中它返回元素的索引值，
	在集合中返回 key-value 对的 key 值。
*/
import (
	"fmt"
)

func main() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	mub := []int{2,3,4};
	//var sum int;
	sum := 0;
	for _,value := range mub {
		sum += value;
	}
	fmt.Printf("sum = %d\n", sum);
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，
	//所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	 for i, num := range mub {
        if num == 3 {
            fmt.Printf("index = %d\n", i);
        }
    }
    //range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"};
    for k, v := range kvs {
    	//map是无须的所以输出不一定按照初始化的顺序打印
        fmt.Printf("%s -> %s\n", k, v);
    }
    //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，
    //第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c);
    }
}