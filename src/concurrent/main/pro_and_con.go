package main

import (
	"fmt"
	"sync"
	"time"
)

type event struct {
	ch      chan int
	buffer  int
	handle  func(num int)
	timeOut time.Duration
}

func NewEvent(buffer int, timeOut time.Duration, handle func(num int)) *event {
	return &(event{
		ch:      make(chan int, buffer),
		handle:  handle,
		timeOut: timeOut,
	})
}

func (e *event) Producer(num int) {
	select {
	case e.ch <- num:
		e.handle(-1)
	case <-time.After(e.timeOut):
	}
}

func (e *event) Consumer() {
	for v := range e.ch {
		e.handle(v)
	}
}
func main() {
	var wg sync.WaitGroup
	event := NewEvent(10, 2*time.Second, func(num int) {
		if num == -1 {
			fmt.Println("生产者生产了产品")
		} else {
			fmt.Println("消费者消费了第", num, "件商品")
		}
	})
	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			event.Producer(i)
		}
		close(event.ch)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		event.Consumer()
		wg.Done()
	}()
	wg.Wait()
}
