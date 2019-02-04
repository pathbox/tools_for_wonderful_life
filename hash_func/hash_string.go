package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var (
	input = flag.String("s", "", "your string")
)

func init() {
	flag.Parse()
}

func main() {
	var err error
	s := *input
	buf := bytes.NewBufferString(s)
	md5Hash := md5.New()
	sha1Hash := sha1.New()
	sha256Hash := sha256.New()
	sha512Hash := sha512.New()

	_, err = md5Hash.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	_, err = sha1Hash.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	_, err = sha256Hash.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	_, err = sha512Hash.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}

	md5Result := md5Hash.Sum(nil)
	sha1Result := sha1Hash.Sum(nil)
	sha256Result := sha256Hash.Sum(nil)
	sha512Result := sha512Hash.Sum(nil)

	fmt.Printf("md5(%s) = %x\n\n", s, md5Result)
	fmt.Printf("sha1(%s) = %x\n\n", s, sha1Result)
	fmt.Printf("sha256(%s) = %x\n\n", s, sha256Result)
	fmt.Printf("sha512(%s) = %x\n\n", s, sha512Result)
	// r := hex.EncodeToString(md5Result)
	// fmt.Println(r)
}

// 输出格式化为16进制后的字符串数据

// hash_str -s=good
