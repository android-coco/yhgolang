package main

import (
	"regexp"
	"fmt"
)

const text = `
My email is ccmouse@gmail.com
email1 is abc@def.org
email2 is kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindString(text)
	//match := re.FindAllString(text,-1)
	//match := re.FindStringSubmatch(text)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
