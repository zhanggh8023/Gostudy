package main

import "fmt"

type Person3 struct {
	name      string
	age       int
	flg       bool
	intereset []string
}

func initFunc(p *Person3) {
	p.name = "nn"
	p.age = 11
	p.flg = true
	p.intereset = append(p.intereset, "shopping")
	p.intereset = append(p.intereset, "sleeping")
}

func initFunc2() *Person3 {
	p := new(Person3)
	p.name = "21vv"
	p.age = 155
	p.flg = true
	p.intereset = append(p.intereset, "shopping2")
	p.intereset = append(p.intereset, "sleeping2")
	return p
}

func main() {
	var person Person3
	initFunc(&person)

	fmt.Println("person:", person)

	p2 := initFunc2()
	fmt.Println("p2:", p2)

}
