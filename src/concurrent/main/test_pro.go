package main

import (
	"fmt"
	"runtime"
	"sync"
)

var sw sync.WaitGroup

func task2(r rune) {
	defer sw.Done()
	for i := 0; i < 1000; i++ {
		fmt.Printf("%c--------->%d\n", r, i)
	}

}
func main() {
	runtime.GOMAXPROCS(2)
	sw.Add(2)
	go task2('A')
	go task2('B')
	sw.Wait()
}
