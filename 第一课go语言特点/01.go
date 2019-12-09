/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/7/10 13:43
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : first_class.go
**/

//Go语言优势劣势分析:1、需要第三方库多样性
//src目录、pkg目录、平台相关目录、bin目录

//Go源码文件：名称以.go为后缀，内容以Go语言代码组织的文件，多个go源码文件是需要用代码包组织起来的
//ds命令与psd命令：
//ds命令的源码文件：goc2p/src/helper/ds/showds.go(——用于显示指定目录的目录结构)
//pds命令的源码文件：goc2p/src/helper/pds/showpds.go(——用于显示指定代码包的依赖关系)
//go build 编译exe
//go install 安装
//go get 下载并安装代码包：
// -d 只执行下载动作，而不执行安装动作；
// -fix 在下载包后先执行修正动作，而后在进行编译和安装；
// -u 更新本地的代码包


package main  // 代码包声明语句。

// 代码包导入语句。
import (
	"fmt"  // 导入代码包fmt。
)

// main函数。
func main() {
	var a int =10
	var p *int =&a

	a=100
	fmt.Println("a=",a)

	*p=250
	fmt.Println("a=",a)
	fmt.Println("*p=",*p)

	// 打印函数调用语句。用于打印输出信息。
	fmt.Println("Go语言编程实战")
}



/*
package main

import (
    "fmt"
)

func main() {
    var (
        num1 int
        num2 int
        num3 int
	)
	num1, num2, num3 = 1, 2, 3
	// 打印函数调用语句。用于打印上述三个变量的值。
	fmt.Println(num1, num2, num3)
}
 */



/*
package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

func (person *Person) Grow() {
	person.Age++
}

//func (person *Person) Move(newAddress string) string {
//	old := person.Address
//	person.Address = newAddress
//	return old
//}

func main() {
	p := Person{"Robert", "Male", 33, "Beijing"}
	p.Grow()
	fmt.Printf("%v\n", p)
}
*/








