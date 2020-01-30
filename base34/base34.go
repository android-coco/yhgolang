package main

import (
	"container/list"
	"errors"
	"fmt"
)

var baseStr = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjklmnpqrstuvwxyz"
var base = []byte(baseStr)
var baseMap map[byte] int


func InitBaseMap(){
	baseMap = make(map[byte]int)
	for i, v := range base {
		baseMap[v] = i
	}
}
func Base34(n uint64)([]byte){
	quotient := n
	mod := uint64(0)
	l := list.New()
	for quotient != 0 {
		//fmt.Println("---quotient:", quotient)
		mod = quotient%34
		quotient = quotient/34
		l.PushFront(base[int(mod)])
		//res = append(res, base[int(mod)])
		//fmt.Printf("---mod:%d, base:%s\n", mod, string(base[int(mod)]))
	}
	listLen := l.Len()

	if listLen >= 6 {
		res := make([]byte,0,listLen)
		for i := l.Front(); i != nil ; i = i.Next(){

			res = append(res, i.Value.(byte))
		}
		return res
	} else {
		res := make([]byte,0,6)
		for i := 0; i < 6; i++ {
			if i < 6-listLen {
				res = append(res, base[0])
			} else {
				res = append(res, l.Front().Value.(byte))
				l.Remove(l.Front())
			}

		}
		return res
	}

}

func Base34ToNum(str []byte)(uint64, error){
	if baseMap == nil {
		return 0, errors.New("no init base map")
	}
	if str == nil || len(str) == 0 {
		return 0, errors.New("parameter is nil or empty")
	}
	var res uint64 = 0
	var r uint64 = 0
	for i:=len(str)-1; i>=0; i-- {
		v, ok := baseMap[str[i]]
		if !ok {
			fmt.Printf("")
			return 0, errors.New("character is not base")
		}
		var b uint64 = 1
		for j:=uint64(0); j<r; j++ {
			b *= 34
		}
		res +=  b*uint64(v)
		r++
	}
	return res, nil
}

func main() {
	InitBaseMap()
	fmt.Printf("len(baseStr):%d, len(base):%d\n", len(baseStr), len(base))
	res := Base34(20)
	fmt.Printf("===============base:24->%s, %d\n", string(res), len(res))
	num, err := Base34ToNum(res)
	if err == nil {
		fmt.Printf("===============base:%s->%d\n", res, num)
	}else {
		fmt.Printf("===============err:%s\n", err.Error())
	}
}
