package main

import "fmt"

type Person2 struct {
	name string
	sex  byte
	age  int
}

func TEST(data string) string {
	test := data
	return test
}

func test2(p *Person2) {
	p.name = "hh"
	p.age = 888
	p.sex = 'm'
}

func main() {
	var p1 *Person2 = &Person2{"n1", 'f', 19}

	fmt.Println("p1=", p1)

	var tmp Person2 = Person2{"n1", 'f', 19}
	var p2 *Person2
	p2 = &tmp
	fmt.Println("p2=", p2)

	test := "hello!"

	fmt.Println("test:", TEST(test))

	p3 := new(Person2)
	p3.name = "mike"
	p3.age = 22
	p3.sex = 'f'
	fmt.Printf("p3,type = %T\n", p3)
	fmt.Println("p3=", p3)

	//fmt.Printf("%p\n",p3)
	//fmt.Printf("%p\n",p3.name)

	test2(p3)
	fmt.Println("p3=", p3)

}
