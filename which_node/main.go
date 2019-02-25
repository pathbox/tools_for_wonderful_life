package main

import (
	"flag"
	"fmt"
	"os"

	json "github.com/json-iterator/go"

	"github.com/parnurzeal/gorequest"
)

type Node struct {
	Name string
	IP   []string
}

type Result struct {
	Node    int
	RetCode int
}

var (
	input = flag.Int("co", 0, "your co ")
)

func main() {
	flag.Parse()
	nodeMap := make(map[int]interface{}, 4)
	node1 := Node{"node1", []string{"192.168.0.1", "192.168.0.1"}}
	node2 := Node{"node2", []string{"192.168.0.1", "192.168.0.1"}}
	node3 := Node{"node3", []string{"192.168.0.13", "192.168.0.1"}}
	node4 := Node{"node4", []string{"192.168.0.1", "192.168.0.1"}}
	nodeMap[1] = map[string]Node{
		"Set1-Node1": node1,
	}
	nodeMap[2] = map[string]Node{
		"Set2-Node2": node2,
	}
	nodeMap[3] = map[string]Node{
		"Set3-Node3": node3,
	}
	nodeMap[4] = map[string]Node{
		"Set4-Node4": node4,
	}

	co := *input

	request := gorequest.New()
	pStr := `{
	"Backend":"UAccountRouter",
	"Action":"GetCompanyNode",
	"CompanyId": %d
	}`

	params := fmt.Sprintf(pStr, co)
	url := "http://example.cn"
	_, body, errs := request.Post(url).
		Set("Content-Type", "application/json").
		Send(params).
		End()

	if errs != nil {
		panic(errs)
		os.Exit(-1)
	}

	result := &Result{}
	err := json.Unmarshal([]byte(body), result)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	m := nodeMap[result.Node]
	fmt.Println()

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(m)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	return

}
