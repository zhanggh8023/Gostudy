/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/7 13:58
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : user.go
**/

package ctrl

import (
	"Gostudy/IM/hello/model"
	"Gostudy/IM/hello/service"
	"Gostudy/IM/hello/util"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func UserLogin(writer http.ResponseWriter, request *http.Request) {
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
		util.RespOk(writer, data, "Success!")
	} else {
		util.RespFail(writer, "密码不正确!")
	}

	//如何返回json
	_, _ = io.WriteString(writer, "hello,world!")
}

var userService service.UserService

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	//mobile,passwd
	_ = request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "Success!")
	}

	//如何返回json
	_, _ = io.WriteString(writer, "hello,world!")
}

func Resp(writer http.ResponseWriter, i int, nil2 interface{}, s string) {

}
