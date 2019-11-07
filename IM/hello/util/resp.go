/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/7 14:13
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : resp.go
**/

package util

import (
	"encoding/json"
	"log"
	"net/http"
)

//逻辑处理
type H struct {
	Code int         `json:"code"` //当返回json首字母大写，可转为小写
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` //*omitempty data为空时不显示
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}
func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
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
