package main

import (
	"fmt"
)

var count = 0

func Count(i int,ch chan int) {
	fmt.Println("Order",i)
	ch <- count
	count++
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(i,chs[i])
	}
	for _, ch := range chs {
		fmt.Println("main", <-ch)
	}
}