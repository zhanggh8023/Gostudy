# Gostudy
一、安装环境
1、下载window10安装包：go1.12.5.windows-amd64.msi，正常解压安装，自动配置好了Path。可以检查一下是否成功：cmd命令：go verison
2、下载goland编辑器：golanddsqi2018.zip，正常破解操作，配置git路径，VCS第二项选择git；github开发者设置中获取Token，goland配置github账户用token登录。
3、完成新建文件头部设置：

* -*- encoding: utf-8 -*-
* @Time    : ${DATE} ${TIME}
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: ${PRODUCT_NAME}
* @File    : ${NAME}.go


4、配置外观显示darcula

二、新建第一条go程序
package main

import "fmt"

//func main() {
//	fmt.Println("Hello, World!")
//}

func main() {
	fmt.Println("This is first go!")

三、编译
目录下命令执行编译：go build .\HelloWorld.go
生成HelloWorld.exe

四、同步更新github代码
