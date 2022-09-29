package main

import (
	"fmt"
	"unsafe"
)

type inter interface {
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	ptr := &arr[1]
	fmt.Println(*ptr) //go语言是不支持直接对指针进行加减法的，也不允许不同类型指针之间的类型转换
	uptr := uintptr(unsafe.Pointer(ptr))
	fmt.Println(uptr)
	uptr += unsafe.Sizeof(arr[1])
	fmt.Println(uptr)
	fmt.Println(*(*int)(unsafe.Pointer(uptr)))
	//fmt.Println(uptr)
	end := unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + unsafe.Sizeof(arr[0]))
	fmt.Println(*(*int)(end))

	i := ^uintptr(0)
	fmt.Printf("%v", i)
}
