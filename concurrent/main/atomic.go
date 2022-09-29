package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

var x int64
var lock2 sync.Mutex

var wg2 sync.WaitGroup

func add() {
	lock2.Lock()
	x++
	lock2.Unlock()
	wg2.Done()
}

func add2() {
	atomic.AddInt64(&x, 1)
	wg2.Done()
}
func main() {
	/*
		start2 := time.Now()
		for i := 0; i < 100000; i++ {
			go add2()
			wg2.Add(1)
		}
		wg2.Wait()
		//end2 := time.Now()
		fmt.Printf("结果为：%d\n花费时间:%f", x, time.Since(start2).Seconds())
	*/
	var i int32 = 10
	atomic.SwapInt32(&i, 2)
	fmt.Println(i)
	var v atomic.Value
	v.Store(&i)
	g := v.Load()
	fmt.Println(g)
	var firstStoreInProgress byte
	fmt.Printf("%d", unsafe.Pointer(&firstStoreInProgress))
	v.Store(nil)
}
