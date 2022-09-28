package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 仓库类型
type event1 struct {
	queue [10]string //存放产品的队列
	num   int        //剩余产品数量
}

var (
	lock sync.Mutex            //锁
	cond = sync.NewCond(&lock) //条件变量
	//产品仓库
	resources = event1{
		num: 0,
	}
	wg sync.WaitGroup
)

// 生产者
func (event *event1) producer(id int) {
	for i := 1; i <= 10; i++ {
		cond.L.Lock()
		for event.num == 10 {
			cond.Wait() //仓库已满，将该goroutine挂起，等待被通知
		}
		//生产产品
		str := strconv.Itoa(id) + "-" + strconv.Itoa(i)
		//将产品装入仓库
		event.queue[event.num] = str
		//计数
		event.num++
		fmt.Println("生产者生产了" + str)
		cond.L.Unlock()
		//该goroutine已解锁，通知在通知等待队列的goroutine获取锁
		cond.Signal()
	}
	wg.Done()
}

func (event *event1) consumer() {
	for i := 1; i <= 10; i++ {
		cond.L.Lock()
		for event.num == 0 {
			cond.Wait()
		}
		event.num--
		fmt.Println("消费者消费了", event.queue[event.num])
		cond.L.Unlock()
		cond.Signal()
	}
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	fmt.Printf("start:%d\n", start)
	//启动10个生产者goroutine
	for i := 1; i <= 10; i++ {
		go resources.producer(i)
		wg.Add(1)
	}
	//启动10个消费者goroutine
	for i := 1; i <= 10; i++ {
		go resources.consumer()
		wg.Add(1)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Printf("end:%d\n", end)
	fmt.Printf("花费时间：%d", end-start)
}
