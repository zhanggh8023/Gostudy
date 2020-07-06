package main

import "fmt"

func main() {
	/*var m1 map[int]string
	//m1[100] = "Green"
	if m1 == nil {
		fmt.Println("map is nil")
	}

	m2 := map[int]string{}
	fmt.Println(len(m2))
	fmt.Println("m2=",m2)
	m2[4]="red"
	fmt.Println("m2=",m2)

	m3 := make(map[int]string) // len
	fmt.Println(len(m3))
	fmt.Println("m3=",m3)
	m3[400]="red"
	fmt.Println("m3=",m3)


	m4 := make(map[int]string,5) // len
	fmt.Println("len=",len(m4))
	//fmt.Println("cap=",cap(m4)) //不能在map 中使用cap
	m4[400]="red"
	fmt.Println("m4=",m4)*/


	//初始化map
	/*var m5 map [int]string = map[int]string{1:"Luffy",2:"Sanli",3:"Xinhua"}
	m6 := map[int]string{1:"Luffy",2:"Sanli",3:"Xinhua"}
	fmt.Println(m5,m6)*/


	m7 := make(map[int]string,10)

	m7[100] = "Nami"
	m7[20] = "Hello"
	m7[3] = "world"

	fmt.Println("m7=",m7)

	m7[3] = "yellow"  //成功！ 将原有map中，key值为3的map元素，替换

	fmt.Println("m7=",m7)

}



