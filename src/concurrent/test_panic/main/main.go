package main

import "fmt"

func task(c chan int) {
	i := <-c
	fmt.Println(i, "已被取出")
}
func main() {
	//var c chan int = nil
	//go task(c)
	//close(c)
	//c <- 1
	//close(c)
	//fmt.Println(i)
	c := make(chan int)
	go task(c)
	c <- 1
	close(c)
	i := <-c
	fmt.Println(i, "被取出")
	fmt.Println("未进入阻塞！")
}
