/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/29 14:59
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : 03.go
**/

package main

import (
	"fmt"
)

func swap(a, b int) {
	a, b = b, a
}

func swap2(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	a, b := 10, 20

	swap(a, b)
	fmt.Println("swap:main: a:", a, "b", b) //调用在函数内成功，没有返回，被销毁

	swap2(&a, &b)
	fmt.Println("swap:main: a:", a, "b", b) //a空间地址赋值b,b空间地址赋值a

}
