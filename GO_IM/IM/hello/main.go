/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/10/12 16:03
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : main.go
**/

package main

import (
	"io"
	"net/http"
)

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login",
		func(writer http.ResponseWriter,
			request *http.Request) {
			//数据库操作
			//逻辑处理
			//restapi json/ xml返回
			_, _ = io.WriteString(writer, "hello,world!")
		})

	//启动web服务器
	_ = http.ListenAndServe(":8085", nil)
}
