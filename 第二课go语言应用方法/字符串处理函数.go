package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "I love my work and I love my family too"

	// 字符串按指定分隔符拆分
	ret := strings.Split(test, " I")
	for _, s := range ret {
		fmt.Println(s)
	}

	// 字符串按空格进行拆分
	ret = strings.Fields(test)
	for _, s := range ret {
		fmt.Println(s)
	}

	// 判断字符串结束标记
	flg := strings.HasSuffix("test.abc", ".abc")
	fmt.Println(flg)

	// 判断字符串起始标记
	flg = strings.HasPrefix("test.abc", "tes")
	fmt.Println(flg)

}
