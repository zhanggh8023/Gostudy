/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/27 9:48
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : contact.go
**/

package model

import "time"

//好友和群都存在这个表里
//可根据具体业务做拆分
type Contact struct {
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//谁的10000,记录是谁的
	Ownerid int64 `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	//对端,10001,对端信息
	Dstobj int64 `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`
	// 什么类型
	Cate int `xorm:"int(11)" form:"cate" json:"cate"`
	// 备注
	Memo string `xorm:"varchar(120)" form:"memo" json:"memo"`
	// 创建时间
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}

const (
	CONCAT_CATE_USER     = 0x01
	CONCAT_CATE_COMUNITY = 0x02
)
