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
	"Gostudy/IM/hello/ctrl"
	"html/template"
	"log"
	"net/http"
)

//解析参数
//如何获得参数
func RegisterView() {
	//解析
	tpl, err := template.ParseGlob("D:/Gostudy/IM/hello/view/**/*")
	//如果报错就不要继续了
	if nil != err {
		//打印并直接退出
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		//fmt.Println(tplname)
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			_ = tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)

	//1、提供静态资源目录支持
	//http.Handle("/",http.FileServer(http.Dir("D:/Gostudy/GO_IM/IM/hello/")))
	//2、提供指定目录的静态文件支持
	//fmt.Println(http.Dir("D:/Gostudy/GO_IM/IM/hello/"))
	http.Handle("/asset/", http.FileServer(http.Dir("D:/Gostudy/IM/hello/")))

	//user/login.shtml
	//http.HandleFunc("/user/login.shtml", func(w http.ResponseWriter, r *http.Request) {
	//	//解析
	//	tpl,err := template.ParseFiles("D:/Gostudy/GO_IM/IM/hello/view/user/login.html")
	//	if nil != err{
	//		//打印并直接退出
	//		log.Fatal(err.Error())
	//	}
	//	tpl.ExecuteTemplate(w,"/user/login.shtml",nil)
	//})
	////注册
	//http.HandleFunc("/user/register.shtml", func(w http.ResponseWriter, r *http.Request) {
	//	//解析
	//	tpl,err := template.ParseFiles("D:/Gostudy/GO_IM/IM/hello/view/user/register.html")
	//	if nil != err{
	//		//打印并直接退出
	//		log.Fatal(err.Error())
	//	}
	//	tpl.ExecuteTemplate(w,"/user/register.shtml",nil)
	//})

	RegisterView()
	//启动web服务器
	_ = http.ListenAndServe(":8085", nil)
}
