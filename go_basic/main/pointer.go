package main

import "fmt"

type animal struct {
	name string
}

func walk(a animal) {
	fmt.Println(a.name + "正在走路！")
}

func main() {
	var p *int
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	i := 1
	p = &i
	a := &animal{"dog"}
	fmt.Println(a)
	p = nil
	fmt.Println(*p)
	animal := animal{name: "cat"}
	walk(animal)
}
