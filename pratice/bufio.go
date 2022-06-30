package pratice

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func ReadString() {
	// https://studygolang.com/articles/21203
	// 缓冲读的大致过程如下，设定好缓冲区大小buf_size后，读取的字节数为rn，缓冲的字节数为bn：
	// 如果缓冲区为空，且 rn >= buf_size，则直接从文件读取，不启用缓冲。
	// 如果缓冲区为空，且 rn < buf_size，则从文件读取buf_size 字节的内容到缓冲区，程序再从缓冲区中读取rn字节的内容，此时缓冲区剩余bn = buf_size - rn字节。
	// 如果缓冲区不为空，rn < bn，则从缓冲区读取rn字节的内容，不发生文件IO。
	// 如果缓冲区不为空，rn >= bn，则从缓冲区读取bn字节的内容，不发生文件IO，缓冲区置为空，回归1/2步骤。

	// 用 strings.Reader 模拟一个文件IO对象
	strReader := strings.NewReader("12345678901234567890123456789012345678901234567890")

	// go 的缓冲区最小为 16 byte，我们用最小值比较容易演示
	bufReader := bufio.NewReaderSize(strReader, 16)
	fmt.Printf("bufReader buffered: %d, \n", bufReader.Buffered())

	// bn = 0 但 rn >= buf_size 缓冲区不启用 发生文件IO
	tmpStr := make([]byte, 16)
	n, _ := bufReader.Read(tmpStr)
	// bufReader buffered: 0, content: 1234567890123456
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 0 rn < buf_size 缓冲区启用
	// 缓冲区从文件读取 buf_size 字节 发生文件IO
	// 程序从缓冲区读取 rn 字节
	// 缓冲区剩余 bn = buf_size - rn 字节
	tmpStr = make([]byte, 15)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 1, content: 789012345678901
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 1 rn > bn
	// 程序从缓冲区读取 bn 字节 缓冲区置空 不发生文件IO
	// 注意这里只能读到一个字节
	tmpStr = make([]byte, 10)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 0, content: 2
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 0 rn < buf_size 启用缓冲读 发生文件IO
	// 缓冲区从文件读取 buf_size 字节
	// 程序从缓冲区读取 rn 字节
	// 缓冲区剩余 bn = buf_size - rn 字节
	tmpStr = make([]byte, 10)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 6, content: 3456789012
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 6 rn <= bn
	// 则程序冲缓冲区读取 rn 字节 不发生文件IO
	tmpStr = make([]byte, 3)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 3, content: 345
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 3 rn <= bn
	// 则程序冲缓冲区读取 rn 字节 不发生文件IO
	tmpStr = make([]byte, 3)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 0, content: 678
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

}

func ReadLine() {
	f, err := os.OpenFile("../tmp/bufio.log", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		logrus.Error("v2ray_accessLog_lastoffset.id open failed")
	}
	defer f.Close()
	line := bufio.NewReaderSize(f, 16)

	for{
		
	}
	lineVal, _, _ := line.ReadLine()
	fmt.Println(string(lineVal))
	a, _ := f.Seek(0, io.SeekCurrent)
	fmt.Println("aaa   ", a)
	lineVal, _, _ = line.ReadLine()
	fmt.Println(string(lineVal))
	a, _ = f.Seek(0, io.SeekCurrent)
	fmt.Println("aaa   ", a)
	lineVal, _, _ = line.ReadLine()
	fmt.Println(string(lineVal))
	a, _ = f.Seek(0, io.SeekCurrent)
	fmt.Println("aaa   ", a)
	lineVal, _, _ = line.ReadLine()
	fmt.Println(string(lineVal))
	a, _ = f.Seek(0, io.SeekCurrent)
	fmt.Println("aaa   ", a)
	lineVal, _, _ = line.ReadLine()
	fmt.Println(string(lineVal))
	a, _ = f.Seek(0, io.SeekCurrent)
	fmt.Println("aaa   ", a)
}
