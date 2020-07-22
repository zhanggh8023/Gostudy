package main

import (
	"fmt"
	"strings"
)

func workCountFunc(str string) map[string]int{
	s :=strings.Fields(str)   //将字符串，拆分成 字符串切片s
	m := make(map[string]int)  // 创建一个用于存储 Word出现次数的map

	// 遍历拆分
	for i := 0; i<len(s);i++{
		if _,ok := m[s[i]] ; ok{    // ok == ture 说明s[i] 这个key存在
			m[s[i]] = m[s[i]]+1   // m[s[i]]++
		}else {					// 说明 s[i] 这个key 不存在，第一次出现。添加到map中
			m[s[i]] = 1
		}
	}
	return m
}

func main() {
	str := "I love my work and I love my family too"
	
	mRet := workCountFunc(str)

	fmt.Println("mRet:", mRet)
}



