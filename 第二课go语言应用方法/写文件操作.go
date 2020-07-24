package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件OpenFile
	t, err := os.OpenFile("D:/Gostudy/第二课go语言应用方法/testFile.xyz", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer t.Close()
	fmt.Println("OpenFile successful")

	// 写入文件
	n, err := t.WriteString("hello ! welcome to GO world !!!\r\n") //win \r\n    linux \n
	if err != nil {
		fmt.Println("Write err:", err)
		return
	}
	fmt.Println("Write successful:", n)

	//写位置
	/*按位置写:
	Seek(): 	修改文件的读写指针位置。
	参1： 偏移量。 正：向文件尾偏， 负：向文件头偏
	参2： 偏移起始位置：
		io.SeekStart: 文件起始位置
		io.SeekCurrent： 文件当前位置
		io.SeekEnd: 文件结尾位置
	返回值：表示从文件起始位置，到当前文件读写指针位置的偏移量。
	off, _ := f.Seek(-5, io.SeekEnd)*/
	v, _ := t.Seek(5, io.SeekEnd)
	fmt.Println("v:", v)

	o, _ := t.WriteAt([]byte("11111"), v)
	fmt.Println("o:", o)

	test := "hello!"

	fmt.Println("test:", test)
}
