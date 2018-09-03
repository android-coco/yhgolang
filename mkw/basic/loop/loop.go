package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"io"
	"strings"
)

func convertToBin(n int)string{
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n >0; n/=2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
func forever()  {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(72387885),
		convertToBin(0),
	)

	printFile("/Users/yhlyl/work/gowork/src/awesomeProject/mkw/abc.txt")

	s :=`fadsfadsf adfadsfasdf fsadfasdf
	fsdafasdf
	fdsafadsfasdf
	 
	
	fd`

	printFileContents(strings.NewReader(s))
	// Uncomment to see it runs forever
	//forever()
}
