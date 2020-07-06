/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/29 15:12
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : 04指针.go
**/

package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //切片【空】，数组指定长度【6】

	//s:=arr[1:3:5]//[low:high:max],对应：起始，结束，容量

	//fmt.Println("s=",s)
	//fmt.Println("len(s)=",len(s))
	//fmt.Println("cap(s)=",cap(s))

	s := arr[1:5:7]
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s))
	fmt.Println("cap(s)=", cap(s)) //未指定跟随原容量10，指定跟随7

	s2 := s[0:6]
	fmt.Println("s=", s2)
	fmt.Println("len(s)=", len(s2))
	fmt.Println("cap(s)=", cap(s2)) //跟随切片容量

}
