package main

import (
	"flag"
	"fmt"

	"github.com/cheynewallace/tabby"
)

func main() {
	flag.Parse()

	t := tabby.New()
	t.AddHeader("Name              ", "Server            ", "Zone    ")
	t.AddLine("server name", "127.0.0.1", "xxxxxx")
	t.AddLine("server name", "127.0.0.1", "xxxxxx")
	t.AddLine("server name", "127.0.0.1", "xxxxxx")
	t.AddLine("server name", "127.0.0.1", "xxxxxx")
	t.AddLine("server name", "127.0.0.1", "xxxxxx")
	t.AddLine("server name", "127.0.0.1", "xxxxxx")
	t.AddLine("server name", "127.0.0.1", "xxxxxx")

	t.Print()
	fmt.Println()

}
