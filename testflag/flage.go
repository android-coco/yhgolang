package main

import (
	"flag"
	"fmt"
)

func main() {
	webserver := flag.Bool("webserver", false, "web server")
	flag.Parse()
	fmt.Println(*webserver)
}
