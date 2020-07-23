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

func main() {
	var p1 *Person2 = &Person2{"n1", 'f', 19}

	fmt.Println("p1=", p1)

	var tmp Person2 = Person2{"n1", 'f', 19}
	var p2 *Person2
	p2 = &tmp
	fmt.Println("p2=", p2)

	test := "hello!"

	fmt.Println("test:", TEST(test))
}
