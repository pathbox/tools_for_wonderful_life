package main

import (
	"flag"
	"fmt"
)

var kgInput = flag.Float64("kg", 0, "kg number")
var lbInput = flag.Float64("lb", 0, "lb number")

func main() {
	flag.Parse()
	kg := *kgInput
	lb := *lbInput
	if kg != 0 {
		lb := 2.2046 * kg
		fmt.Printf("%vlb\n", lb)
	} else if lb != 0 {
		kg := lb / 2.2046
		fmt.Printf("%vkg\n", kg)
	} else {
		fmt.Println(0)
	}
	return
}
