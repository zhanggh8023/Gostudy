package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	p     []Person
	id    []int
	score [1000]int
}

func test(man Person) {
	fmt.Println("test temp size:", unsafe.Sizeof(man))

	man.name = "name1"
	man.age = 33

	fmt.Println("test:man", man)

}

func main() {
	// 顺序初始化
	var man Person = Person{"andy", 'm', 18}

	fmt.Println("man", man)

	// 部分初始化
	man2 := Person{sex: 'f', age: 19}
	fmt.Println("man2", man2)

	// 索引成员变量
	fmt.Printf("man2.name =%q\n", man2.name)

	var man3 Person
	man3.name = "mike"
	man3.sex = 'm'
	man3.age = 99
	fmt.Println("man3", man3)
	man3.age = 9990
	fmt.Println("man3", man3)

	// 结构体比较
	var p1 Person = Person{"andy", 'm', 18}
	var p2 Person = Person{"andy", 'm', 18}
	var p3 Person = Person{"andy", 'm', 181}

	fmt.Println("p1==p2", p1 == p2)
	fmt.Println("p1==p3", p1 == p3)

	// 相同类型结构体赋值
	var tmp Person
	fmt.Println("tmp", tmp)

	tmp = p3
	fmt.Println("tmp", tmp)

	// 函数内部使用结构图传参
	var temp Person
	fmt.Println("temp:", unsafe.Sizeof(temp))
	test(temp)
	fmt.Println("temp", temp)

	fmt.Printf("&temp = %p\n", &temp)
	fmt.Printf("&temp.name = %p\n", &temp.name)

	fmt.Println("temp:", unsafe.Sizeof(true))

}
