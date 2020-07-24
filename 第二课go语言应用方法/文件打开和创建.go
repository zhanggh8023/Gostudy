package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件
	f, err := os.Create("D:/Gostudy/第二课go语言应用方法/testFile.xyz")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}

	defer f.Close()
	fmt.Println("Create successful")

	/*// 打开文件OpenFile
	t, err := os.OpenFile("D:/Gostudy/第二课go语言应用方法/testFile.xyz", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer t.Close()
	fmt.Println("OpenFile successful")

	// 写入文件
	_, err = t.WriteString("hello! welcome to GO world!")
	if err != nil {
		fmt.Println("Write err:", err)
		return
	}
	defer t.Close()
	fmt.Println("Write successful")


	// 打开文件Open
	m, err := os.Open("D:/Gostudy/第二课go语言应用方法/testFile.xyz")
	if err != nil {
		fmt.Println("Open err:", err)
		return
	}

	//defer m.Close()
	fmt.Println("Open successful")

	// 读文件
	i, err := m.Read([]byte("D:/Gostudy/第二课go语言应用方法/testFile.xyz"))
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}

	fmt.Println("OpenFile successful:", i)*/

	test := "hello!"

	fmt.Println("test:", test)
}
