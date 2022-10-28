package main

import (
	"fmt"
	"time"
)

type notifier interface {
	SendInt(c chan<- string) bool
}
type send struct {
	message string
}

func (s send) SendInt(c chan<- string) bool {
	c <- s.message
	return true
}
func task(c chan int) {
	time.Sleep(2 * time.Second)
	c <- 1
}

func main() {
	c := make(chan int)
	go task(c)
	//获取一个通道，并且指定获取值需要等待的时间
	timeOut := time.After(3 * time.Second)
	select {
	case i := <-c:
		fmt.Println("已获取i:", i)
	case <-timeOut:
		fmt.Println("超时未获取成功！")
	default:
		fmt.Println("默认分支执行")
	}
	chan1 := make(chan int, 2)
	chan2 := make(chan int, 2)
	close(chan1)
	//go task(chan2)
	select {
	case i, ok := <-chan1:
		if !ok {
			chan1 = nil
		} else {
			fmt.Println("chan1取出数据", i)
		}
	case chan2 <- 1:
		fmt.Println("chan2获取数据")
	}
}
