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
