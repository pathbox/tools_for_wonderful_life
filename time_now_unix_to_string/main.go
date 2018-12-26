package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

var (
	tp = flag.Int64("tp", time.Now().Unix(), "timestamp value")
)

func main() {
	flag.Parse()
	t := *tp
	r := timeNowUnixToString(t)
	fmt.Println(r)
}

func timeNowUnixToString(timestamp int64) string {
	if timestamp == 0 {
		n := time.Now().Unix()
		return strconv.Itoa(int(n))
	}
	tm := time.Unix(timestamp, 0)
	// fmt.Println(tm.Format("2006-01-02 03:04:05 PM")) // 两种格式化模式
	// fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))
	s := tm.Format("2006-01-02 15:04:05")
	return s
}

// go run main.go 获得当前now的时间
// go run main.go -tp=1 获得指定时间戳的时间表达式
