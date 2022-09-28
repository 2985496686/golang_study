package main

import (
	"fmt"
	"sync"
	time "time"
)

var (
	i     int
	group sync.WaitGroup
	rw    sync.RWMutex
)

func read() {
	rw.RLock()
	time.Sleep(time.Millisecond)
	rw.RUnlock()
	group.Done()
}

func write() {
	rw.Lock()
	i++
	rw.Unlock()
	group.Done()
}
func main() {
	fmt.Println("开始执行：", time.Now(), "\n")
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go read()
	}
	group.Wait()
	//time.Sleep(3 * time.Second)
	fmt.Printf("执行结束：%v   i = %d", time.Now(), i)
	//rw.RUnlock()
	var rwLock sync.RWMutex
	rwLock.Lock()
	rwLock.Unlock()
	rwLock.RLock()
	rwLock.RUnlock()
}
