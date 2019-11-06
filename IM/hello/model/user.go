/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/5 15:05
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : user.go
**/

package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MAN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	Id       int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`  //用户ID
	Mobile   string    `xorm:"varchar(20)" form:"mobile" json:"mobile"`     //手机号码
	Passwd   string    `xorm:"varchar(40)" form:"passwd" json:"-"`          // 用户密码=f(plainpwd+salt),MD5
	Avatar   string    `xorm:"varchar(150)" form:"avatar" json:"avatar"`    //头像
	Sex      string    `xorm:"varchar(2)" form:"sex" json:"sex"`            // 性别
	Nickname string    `xorm:"varchar(20)" form:"nickname" json:"nickname"` // 昵称
	Salt     string    `xorm:"varchar(10)" form:"salt" json:"-"`            // 随机数
	Online   int       `xorm:"int(10)" form:"online" json:"online"`         //是否在线
	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`       // /chat?id=1&token=x
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`        // 备注
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`    // 统计每天增加的用户
}
