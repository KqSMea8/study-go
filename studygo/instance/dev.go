//
// 这是关于单例中 用锁的解法
//
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


type mytest struct {
	name string
}

var myInstance *mytest
var mu sync.Mutex	//加锁包
var initNum uint32

//单例 非线程
//实现 加锁
//带检查锁的
//设置标记
func instance(tim int) *mytest {
	if myInstance==nil {	//带检查锁的
		if atomic.LoadUint32(&initNum) == 1 {
			fmt.Println("mmyInstance is nil,atomic:", " value:", myInstance)
			return myInstance
		}
		mu.Lock()
		defer mu.Unlock()
		if myInstance == nil {
			myInstance = &mytest{
				name:"xiaojh",
			}
			atomic.StoreUint32(&initNum,1)
			fmt.Println("mmyInstance is nil,tim:", tim, " value:", myInstance)
		} else {
			fmt.Println("mmyInstance is setting,tim:", tim, " value:", myInstance)
		}
	}
	return myInstance
}

func main() {
	for i:=0;i<4;i++ {
		instance(i)
	}

	var input string
	fmt.Scanln(&input)	//这里的 Scanln 代码需要我们在程序退出前按下任意键结束。
	fmt.Println("done")

}
