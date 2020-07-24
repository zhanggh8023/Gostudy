package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	test := "hello!"
	fmt.Println("test:", test)

	// 打开读文件
	f_r, err := os.Open("D:/Gostudy/第二课go语言应用方法/testFile.xyz")
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f_r.Close()
	fmt.Println("OpenFile successful:", f_r)

	// 创建写文件
	f_w, err := os.Create("D:/Gostudy/第二课go语言应用方法/testw.txt")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f_w.Close()
	fmt.Println("Create successful:", f_w)

	// 从读文件中获取数据，放到缓冲区
	buf := make([]byte, 4096)
	// 循环从读文件中，获取数据
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("reading n = %d\n", n)
			return
		}
		f_w.Write(buf[:n])
		fmt.Printf("reading n = %d\n", n)
	}

}
