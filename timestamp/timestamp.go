package main

import (
	"flag"
	"fmt"
	"time"
)

var opInput = flag.String("op", "+", "+/-")
var timeInput = flag.String("time", "day", "time unit")
var numInpout = flag.Int("num", 0, "number")

func main() {
	flag.Parse()
	num := *numInpout
	if num == 0 {
		ts := time.Now().Unix()
		fmt.Println(ts)
		return
	}
}
