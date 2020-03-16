/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/29 14:17
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : 02.go
**/

package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println("&a=", &a)

	var p *int //默认类型的 默认值

	p = new(int)
	//p=new(string)

	*p = 1000 //左值表示获取空间，写
	m := *p   //右值表示获取值，读

	fmt.Println("变量：", m)
	fmt.Printf("%s\n", *p)
	fmt.Printf("%q\n", *p) //打印GO语言格式的字符串

}
