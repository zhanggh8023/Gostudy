package main

import "fmt"

//拷贝
/*func main()  {
	data := [...]int{0,1,2,3,4,5,6,7,8,9}
	s1 := data[8:]  //【8，9】
	s2 := data[0:5]  // 【0，1，2，3，4】

	copy(s2,s1) //【拷贝区】【拷贝对象】，对应拷贝去位置【8，9，2，3，4】
	fmt.Println("s2=",s2)

}*/

//去除
func remove(data []int,idx int) []int  {
	copy(data[idx:],data[idx+1:])
	return data[:len(data)-1]

}


func main()  {
	data := []int{5,6,7,8,9}

	afterData := remove(data,2)

	fmt.Println("afterData=",afterData)

}




