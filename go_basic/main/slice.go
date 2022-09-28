package main

import "fmt"

func remake(arr *[]int) {
	*arr = (*arr)[0:2]

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%d  %d\n", len(arr), cap(arr)) //6  6
	remake(&arr)
	fmt.Printf("%d  %d", len(arr), cap(arr)) //2  6
}
