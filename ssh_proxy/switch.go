package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	c = flag.String("c", "y", "y: 开启 n:关闭")
)

func main() {
	flag.Parse()
	str := *c
	var err error
	cmd := exec.Command("sh", "-c", "export http_proxy=http://127.0.0.1:35410")
	if str != "y" {
		fmt.Println("Switch switch my teminal network proxy off")
		cmd = exec.Command("sh", "-c", "export http_proxy=")
	} else {
		fmt.Println("switch my teminal network proxy on")
	}
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("switch my teminal network proxy success~")
}

// export http_proxy=http://127.0.0.1:35410;

func Cmd(commandName string, params []string) (string, error) {
	cmd := exec.Command(commandName, params...)
	fmt.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}
