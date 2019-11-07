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
	"Gostudy/IM/hello/model"
	"Gostudy/IM/hello/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex

	//获取随机数，进行md5加密
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpsd, tmp.Salt)
	tmp.Createat = time.Now() //统计用户总数
	//token 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	//返回新用户信息

	//插入InsertOne
	_, err = DbEngin.InsertOne(&tmp)
	//前端恶意插入特殊字符
	//数据库连接操作失败
	return tmp, nil
}

//登录函数
func (s *UserService) Login(
	mobile, //手机
	plainpsd string) (user model.User, err error) {
	return user, nil
}
