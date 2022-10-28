package main

import (
	"runtime"
	"sync/atomic"
)

type spinlock int32

func (sl *spinlock) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(sl), 0, 1) {
		runtime.Gosched() //让出CPU时间片
		continue
	}
}

func (sl *spinlock) Unlock() {
	if atomic.LoadInt32((*int32)(sl)) == 0 {
		panic("error,unlock a unlocked lock")
	} else {
		atomic.StoreInt32((*int32)(sl), 0)
	}
}

func main() {
	var lock spinlock
	lock.Lock()
	lock.Lock()
	lock.Unlock()
}
