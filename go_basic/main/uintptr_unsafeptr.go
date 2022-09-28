package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	ptr := &arr[0]
	fmt.Println(*ptr) //go语言是不支持直接对指针进行加减法的，也不允许不同类型指针之间的类型转换
}
