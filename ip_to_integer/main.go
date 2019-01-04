package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	input = flag.String("ip", "0.0.0.0", "input IP")
)

func main() {
	flag.Parse()
	ip := *input
	inputIP(ip)
}

func inputIP(ip string) {
	x := strings.Split(ip, ".")
	b0, _ := strconv.ParseInt(x[0], 10, 0) //字符串转Int64型
	b1, _ := strconv.ParseInt(x[1], 10, 0)
	b2, _ := strconv.ParseInt(x[2], 10, 0)
	b3, _ := strconv.ParseInt(x[3], 10, 0)
	getResult(b0, b1, b2, b3)
}

func getResult(b0, b1, b2, b3 int64) {
	/*
	 *用于将IP地址转换成10进制
	 *需要分别输入IP地址的4段内容
	 *由于计算之后的数字很大所以需要用到int64类型
	 *分别计算IP地址的4段；然后相加即可
	 */

	num0 := b0 * 16777216 // 256*256*256 最高位,每升一位，其实就是256倍
	num1 := b1 * 65536    // 256*256
	num2 := b2 * 256
	num3 := b3 * 1
	sum := num0 + num1 + num2 + num3
	fmt.Println("转换之后的十进制IP地址值:", sum)
}

// go run main.go -ip="127.0.0.1"
