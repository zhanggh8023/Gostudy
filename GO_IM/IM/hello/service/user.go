/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/5 17:38
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : user.go
**/

package service

import (
	"Gostudy/GO_IM/IM/hello/model"
	"errors"
)

type UserService struct {
}

//注册函数
func (s *UserService) Register(
	mobile, //手机
	plainpsd, //明文密码
	nickname, //昵称
	avatar, sex string) (user model.User, err error) {
	//检测手机号码是否存在
	tmp := model.User{}
	_, err = DbEngin.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	//如果存在则返回提示已注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册！")
	}
	//否则拼接插入数据

	//返回新用户信息

	return user, nil
}

//登录函数
func (s *UserService) Login(
	mobile, //手机
	plainpsd string) (user model.User, err error) {
	return user, nil
}
