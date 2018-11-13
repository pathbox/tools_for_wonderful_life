package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type sysFileInfo struct {
	Name  string
	Size  int64
	Mtime time.Time
	Perm  os.FileMode
	Md5   string
	Type  bool
}

var (
	listenPort = flag.String("port", "9696", "server listen port")
	syncFile   = flag.String("file", "", "transfer file")
	syncHost   = flag.String("host", "", "server host")
	syncSer    = flag.Bool("d", false, "server mode") // true:服务端接收文件，false: 客户端发送文件
	syncFold   = flag.String("dir", "/tmp/filesync/", "recive sync fold ")
)

func main() {
	flag.Parse() // 第一步要解析 flag
	if *syncSer {
		servPort := fmt.Sprintf(":%s", *listenPort)
		l, err := net.Listen("tcp", servPort)
		if err != nil {
			fmt.Println("net listen failed:", err)
			return
		}
		err = os.MkdirAll(*syncFold, 0755)
		if err != nil {
			panic(err)
		}
		fmt.Println("Start Service")
		Serve(l) // 开始服务，接收请求连接
	} else {
		destination := fmt.Sprintf("%s:%s", *syncHost, *listenPort)
		clientSend(*syncFile, destination)
	}
}

func clientSend(files string, destination string) {
	fInfo := getFileInfo(files)
	newName := fInfo.Name
	cmdLine := fmt.Sprintf("upload %s %d %d %d %s ", newName, fInfo.Mtime.Unix(), fInfo.Perm, fInfo.Size, fInfo.Md5)
	client, err := net.Dial("tcp", destination)
	if err != nil {
		fmt.Println("client connect error:", err)
		return
	}
	defer client.Close()

	client.Write([]byte(cmdLine))
	client.Write([]byte("\r\n"))
	fileHandle, err := os.Open(files)
	if err != nil {
		fmt.Println("open file ERROR", err)
		return
	}

	io.Copy(client, fileHandle)

	for {
		buffer := make([]byte, 4096)    // 4k
		num, err := client.Read(buffer) // 每次读取4k
		if err == nil && num > 0 {
			fmt.Println(string(buffer[:num]))
			break
		}
	}
}

func getFileInfo(filename string) *sysFileInfo {
	f, err := os.Lstat(filename)
	if err != nil {
		fmt.Println("info ERROR", err)
		return nil
	}

	fileHandle, err := os.Open(filename)
	if err != nil {
		fmt.Println("open ERROR", err)
		return nil
	}

	h := md5.New()
	_, err = io.Copy(h, fileHandle)
	fileInfo := &sysFileInfo{
		Name:  f.Name(),
		Size:  f.Size(),
		Perm:  f.Mode().Perm(),
		Mtime: f.ModTime(),
		Type:  f.IsDir(),
		Md5:   fmt.Sprintf("%x", h.Sum(nil)),
	}
	return fileInfo
}

func Serve(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				continue
			}
			fmt.Println("network error", err)
		}
		go Handler(conn)
	}
}

func Handler(conn net.Conn) {
	defer conn.Close()

	state := 0
	var cmd *sysFileInfo
	var fSize int64
	var tempFileName string
	var n int64
	for {
		buffer := make([]byte, 2048)
		num, err := conn.Read(buffer)
		numLen := int64(num)
		if err != nil && err != io.EOF {
			fmt.Println("cannot read", err)
		}
		n = 0
		if state == 0 {
			n, cmd = cmdParse(buffer[:num])
			tempFileName = fmt.Sprintf("%s.newsync", cmd.Name)
			fSize = cmd.Size
			state = 1
		}
		if state == 1 {
			last := numLen
			if fSize <= numLen-n {
				last = fSize + n
				state = 2
			}
			err = writeToFile(buffer[int(n):int(last)], tempFileName, cmd.Perm)
			if err != nil {
				fmt.Println("read num error : ", err)
			}
			fSize -= last - n
			if state == 2 {
				os.Remove(cmd.Name)
				err = os.Rename(tempFileName, cmd.Name)
				if err != nil {
					fmt.Println("rename ", tempFileName, " to ", cmd.Name, " failed")
				}
				err = os.Chtimes(cmd.Name, time.Now(), cmd.Mtime)
				if err != nil {
					fmt.Println("change the mtime error ", err)
				}
				fileHandle, err := os.Open(cmd.Name)
				if err != nil {
					fmt.Println("open ERROR", err)
				}
				h := md5.New()
				io.Copy(h, fileHandle)
				newfMd5 := fmt.Sprintf("%x", h.Sum(nil))
				if newfMd5 == cmd.Md5 {
					sendInfo := fmt.Sprintf("%s sync success", cmd.Name)
					conn.Write([]byte(sendInfo))
				} else {
					sendInfo := fmt.Sprintf("%s sync failed", cmd.Name)
					conn.Write([]byte(sendInfo))
				}
			}
		}
	}
}

func cmdParse(infor []byte) (int64, *sysFileInfo) {
	var i int64
	for i = 0; i < int64(len(infor)); i++ {
		if infor[i] == '\n' && infor[i-1] == '\r' {
			cmdLine := strings.Split(string(infor[:i-1]), " ")
			fileName := fmt.Sprintf("%s/%s", *syncFold, cmdLine[1])
			filePerm, _ := strconv.Atoi(cmdLine[3])
			fileMtime, _ := strconv.ParseInt(cmdLine[2], 10, 64)
			fileSize, _ := strconv.ParseInt(cmdLine[4], 10, 64)
			fileInfo := &sysFileInfo{
				Name:  fileName,
				Mtime: time.Unix(fileMtime, 0),
				Perm:  os.FileMode(filePerm),
				Size:  fileSize,
				Md5:   string(cmdLine[5]),
			}
			return i + 1, fileInfo
		}
	}
	return 0, nil
}

func writeToFile(data []byte, fileName string, perm os.FileMode) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, perm)
	if err != nil {
		fmt.Println("write file error:", err)
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("write file error", err)
		return err
	}
	return nil
}

/*
Server:
Usage of ./gosync:
  -d=false: server mode
  -dir="/tmp/gosync/": recive sync fold
  -file="": transfer file
  -host="": server host
	-port="7722": server listen port

	./gosync -d; -dir; /tmp/xxxx

Client:


./gosync --file=/root/Videos/DSCF2394.avi --host=127.0.0.1 --port=7722

# md5sum /tmp/xxxx/DSCF2394.avi
eb50332d3b3b6f36b773046aca16e908  /tmp/xxxx/DSCF2394.avi
# md5sum /root/Videos/DSCF2394.avi
eb50332d3b3b6f36b773046aca16e908  /root/Videos/DSCF2394.
*/
