package main

import (
	"flag"
	"fmt"
)

var (
	s = flag.String("s", "", "your string")
)

func main() {
	flag.Parse()
	str := *s
	n := len(str)
	fmt.Println(n)
}

// go run main.go -s="hallo"
