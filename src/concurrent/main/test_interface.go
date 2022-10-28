package main

import (
	"fmt"
	"reflect"
)

type animal interface {
	move()
}

type cat struct {
}

/*
func (c cat) move() {
	fmt.Printf("cat正在移动")
}*/

func (c *cat) move() {
	fmt.Printf("cat正在移动")
}

func mySelect(x interface{}) {
	switch i := x.(type) {
	case int32:
		fmt.Printf("类型是：%c\n", i)
	}
}
func main() {
	type cat struct {
	}
	mySelect('s')
	s := "ss"
	typ := reflect.TypeOf(&s)
	typOfCatPtr := reflect.TypeOf(&cat{})
	typOfCat := typOfCatPtr.Elem()
	fmt.Println(typ.Kind())
	fmt.Println(typOfCatPtr.Name())
	fmt.Println(typOfCat.Name())
}
