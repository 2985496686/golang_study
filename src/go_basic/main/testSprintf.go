package main

import "fmt"

type User struct {
	Name string
	pet
}
type pet struct {
	petName string
}

func (p pet) move() {
	fmt.Println(p.petName + "正在移动！")
}

func (user User) String() string {
	return "user:" + user.Name
}
func main() {
	p := User{
		Name: "张三",
		pet:  pet{"小狗"},
	}
	p.move()
	fmt.Println(p)
}
