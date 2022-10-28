package main

import "fmt"

type Walker interface {
	Walk()
}

type user2 struct {
	name string
	pet2
}

type pet2 struct {
	pName string
}

func (p pet2) Walk() {
	fmt.Println(p.pName + "正在移动")
}
func SuperWalk(walker Walker) {
	walker.Walk()
	fmt.Println("非常快！")
}
func main() {
	p := user2{
		name: "张三",
		pet2: pet2{pName: "小狗"},
	}

	SuperWalk(p)
}
