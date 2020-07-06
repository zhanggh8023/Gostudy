/**
* -*- encoding: utf-8 -*-
* @Time    : 2019/11/29 16:42
* @Author  : zgh
* @Email   : 849080458@qq.com
* @Software: GoLand
* @File    : 05切片的使用.go
**/

package main

import "fmt"

//创建切片
/*func main(){
	arr := [8]int{1,2,3,4,5,6,7,8}

	//s := arr[1:3:5]
	//
	//fmt.Println("s=",s)
	//fmt.Println("len(s)",len(s))
	//fmt.Println("cap(s)",cap(s))

	s := arr[1:5]

	fmt.Println("s=",s)
	fmt.Println("len(s)",len(s))
	fmt.Println("cap(s)",cap(s))
}*/

/*func main()  {
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	s := arr[2:5]
	fmt.Println("s=",s)
	fmt.Println("len(s)",len(s))
	fmt.Println("cap(s)",cap(s))

	s2 := arr[2:7]
	fmt.Println("s=",s2)
	fmt.Println("len(s)",len(s2))
	fmt.Println("cap(s)",cap(s2))
}*/



/*func main() {
	arr := [10]int {1,2,3,4,5,6,7,8,9,10}

	s:= arr[1:5:7]


	fmt.Println("s=",s)
	fmt.Println("len(s）=", len(s))  //5-1=4
	fmt.Println("cap(s)=", cap(s))  //7-1


	s2:=s[0:6]
	fmt.Println("s=",s2)
	fmt.Println("len(s）=", len(s2))  //6-0
	fmt.Println("cap(s)=", cap(s2))   //
}*/


//切片使用
/*func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := arr[2:5:5]

	fmt.Println("s=", s)
	fmt.Println("len(s）=", len(s))
	fmt.Println("cap(s)=", cap(s))

	s2 := s[2:7]

	fmt.Println("s2=", s2)
	fmt.Println("len(s2）=", len(s2))
	fmt.Println("cap(s2)=", cap(s2))

}*/

//切片创建
/*func main(){
	//自动推导赋初值
	s1 := []int {1,2,4,6}
	fmt.Println("s1=",s1)

	s2 := make([]int,5,10)
	fmt.Println("len=",len(s2),"cap=",cap(s2))

	s3 := make([]int,7)
	fmt.Println("len=",len(s3),"cap=",cap(s3))
}*/

/*func main()  {
	s1 :=  []int {1,2,4,6}

	s1 = append(s1,888)
	s1 = append(s1,888)
	s1 = append(s1,888)
	s1 = append(s1,888)
	s1 = append(s1,888)
	s1 = append(s1,888)

	fmt.Println("s1=",s1)
}*/

//练习
func noEmpty(data []string) []string {
	out := data[:0]
	for _, str := range data{
		if str !="" {
			out =append(out,str)
		}
		// 去空
	}
	return out
}

// 原值处理
func noEmpty2(data []string) []string{
	i := 0
	for _, str := range data{
		if str !="" {
			data[i]=str
			i++
		}
		// 去空
	}
	return data[:i]
}

// 消除重复字符串
func noSame(data []string) []string {
	out := data[:1]
	for _, str := range data {
		i:=0
		for ;i<len(out);i++ {
			if str == out[i]{
				break
			}
		}
		if i == len(out){
			out = append(out,str)
		}

	}
	return out
}




func main()  {
	// {"red" "black", "pink", "", "", "pink", "blue"}
	// {"red", "", "black", "", "", "pink", "blue"}
	// {"red","black", "red","pink", "blue","pink", "blue"}
	data := []string{"red","black", "red","pink", "blue","pink", "blue"}
	afterData := noSame(data)
	fmt.Println("afterData:",afterData)
}










