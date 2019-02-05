package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	json "github.com/json-iterator/go"
)

var (
	input = flag.String("no", "", "your string")
	k     = flag.String("k", "", "快递公司类型")
)

type Data struct {
	Time     string
	FTime    string `json:"ftime"`
	Context  string `json:"context"`
	Location string `json:"location"`
}
type Message struct {
	Message   string
	Nu        string
	IsCheck   string
	Condition string
	Com       string
	Status    string
	State     string
	Data      []Data
}

func init() {
	flag.Parse()
}

func main() {
	no := *input
	url := fmt.Sprintf("http://www.kuaidi100.com/query?type=%s&postid=%s", *k, no)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	m := &Message{}

	err = json.Unmarshal(buf, m)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	fmt.Println(m)
}
