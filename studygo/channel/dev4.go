package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func main() {
	var pp int
	for i := 1; i < 101; i++ {
		w.Add(1)
		go func(n int) {
			pp = pp+n
			fmt.Println(n,pp)
			w.Done()
		}(i)
	}
	// 等待一百个协程完成。
	w.Wait()
	fmt.Println("test:",pp)
}