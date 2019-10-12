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
			//1、获取前端传递的参数
			//mobile,passwd
			//解析参数
			//如何获得参数
			_ = request.ParseForm()
			mobile := request.PostForm.Get("mobile")
			passwd := request.PostForm.Get("passwd")

			loginok := false
			if mobile == "18600000000" && passwd == "123456" {
				loginok = true
			}
			str := `{"code":0, "data":{"id":1, "token":"test"}}`
			if !loginok {
				str = `{"code":-1, "msg":"密码不正确！"}`
				//返回错误json
			}
			//设置header 为json  默认的text/html，所以特别指出返回为application/json
			writer.Header().Set("Content-Type", "application/json")
			//设置状态200
			writer.WriteHeader(http.StatusOK)
			//输出
			_, _ = writer.Write([]byte(str))
			//返回json ok
			//如何返回json
			_, _ = io.WriteString(writer, "hello,world!")
		})

	//启动web服务器
	_ = http.ListenAndServe(":8085", nil)
}
