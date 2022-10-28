package main

import (
	"fmt"
)

func task(c chan int) {
	close(c)
}

func main() {
	c := make(chan int, 4)

	c <- 1
	c <- 2
	close(c)
	for i := range c {
		fmt.Println(i)
	}
}
