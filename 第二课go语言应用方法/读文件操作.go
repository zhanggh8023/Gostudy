package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	test := "hello!"
	fmt.Println("test:", test)

	// 打开文件OpenFile
	t, err := os.OpenFile("D:/Gostudy/第二课go语言应用方法/testFile.xyz", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer t.Close()
	fmt.Println("OpenFile successful")

	// 创建一个带有缓冲区的 reader
	reader := bufio.NewReader(t)
	for {
		buf, err := reader.ReadBytes('\n') //读一行
		if err != nil && err == io.EOF {
			fmt.Println("文件读取完成！")
			return
		} else if err != nil {
			fmt.Println("ReadBytes err:", err)
		}
		fmt.Println("ReadBytes successful:", string(buf))
	}

}
