package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error")
	err2 := err
	fmt.Printf("%p,%p", err, err2)
	fmt.Println(err)
	//panic("sssssss")
}
