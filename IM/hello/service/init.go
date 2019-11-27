/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/5 17:51
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : init.go
**/
package service

import (
	"Gostudy/IM/hello/model"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	//下划线:当导入一个包时，该包下的文件里所有init函数都会被执行
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//数据库操作
//安装工具：go get github.com/go-xorm/xorm
//安装驱动：go get github.com/go-sql-driver/mysql
var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	//DsName := "root:****@(47.110.127.36:3306)/chat?charset=utf8"
	DsName := "root:123456@(172.16.20.130:3306)/chat?charset=utf8"

	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)

	if err != nil {
		log.Fatal(err.Error())
	}

	//是否显示SQL语句
	DbEngin.ShowSQL(true)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)

	//自动创建表结构User
	_ = DbEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community))
	//DbEngin := dbengin
	fmt.Println("init data base ok!")
}
