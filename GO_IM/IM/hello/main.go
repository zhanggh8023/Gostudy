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
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

func userLogin(writer http.ResponseWriter, request *http.Request) {

	//mobile,passwd
	_ = request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	loginok := false
	if mobile == "18600000000" && passwd == "123456" {
		loginok = true
	}
	if loginok {
		//｛"id":1,"token":"xx"｝
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer, 0, data, "Success!")
	} else {
		Resp(writer, -1, nil, "密码不正确!")
	}

	//如何返回json
	_, _ = io.WriteString(writer, "hello,world！")
}

//逻辑处理
type H struct {
	Code int         `json:"code"` //当返回json首字母大写，可转为小写
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` //*omitempty data为空时不显示
}

//restapi json/ xml返回
//1、获取前端传递的参数
func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	//设置header 为json  默认的text/html，所以特别指出返回为application/json
	w.Header().Set("Content-Type", "application/json")
	//设置状态200
	w.WriteHeader(http.StatusOK)
	//输出
	//定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//将结构体转化成json字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	_, _ = w.Write(ret)
	//返回json ok
}

//解析参数
//如何获得参数
func RegisterView() {
	//解析
	tpl, err := template.ParseGlob("D:/Gostudy/GO_IM/IM/hello/view/**/*")
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
	http.HandleFunc("/user/login", userLogin)

	//1、提供静态资源目录支持
	//http.Handle("/",http.FileServer(http.Dir("D:/Gostudy/GO_IM/IM/hello/")))
	//2、提供指定目录的静态文件支持
	//fmt.Println(http.Dir("D:/Gostudy/GO_IM/IM/hello/"))
	http.Handle("/asset/", http.FileServer(http.Dir("D:/Gostudy/GO_IM/IM/hello/")))

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
