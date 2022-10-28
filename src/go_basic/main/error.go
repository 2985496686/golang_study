package main

import (
	"errors"
	"fmt"
	ioutil "io/ioutil"
	"os"
)

func checkIndex(i int, arr []string) error {
	if i < 0 || i >= len(arr) {
		return errors.New("error:" + "out of bound")
		panic("出错！")
	} else {
		return nil
	}
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	defer func() {
		recover()
	}()
	arr := []string{"1", "2"}
	fmt.Println(checkIndex(-1, arr))
}
