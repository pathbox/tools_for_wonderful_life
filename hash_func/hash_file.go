package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

var (
	input = flag.String("f", "", "your file path")
)

func init() {
	flag.Parse()
}

func main() {
	var err error
	f := *input

	st := time.Now()

	data, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(data)
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

	md5Str := hex.EncodeToString(md5Result)
	sha1Str := hex.EncodeToString(sha1Result)
	sha256Str := hex.EncodeToString(sha256Result)
	sha512Str := hex.EncodeToString(sha512Result)

	fmt.Printf("File Size: %dK, time: %s\n\n", len(data)/1000, time.Since(st))
	fmt.Printf("md5(%s) = len(%d) %s\n\n", f, len(md5Str), md5Str)
	fmt.Printf("sha1(%s) = len(%d) %s\n\n", f, len(sha1Str), sha1Str)
	fmt.Printf("sha256(%s) = len(%d) %s\n\n", f, len(sha256Str), sha256Str)
	fmt.Printf("sha512(%s) = len(%d) %s\n\n", f, len(sha512Str), sha512Str)

}

// hash_file -f="path/file.txt"
