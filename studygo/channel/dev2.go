// 对struct 进行加锁
// 同一变量如果处理
//
package main

import (
	"fmt"
	"sync"
)

var ll sync.Mutex
var aa string

func f() {
	aa = "hello world"
	fmt.Println(aa)
	//ll.Unlock()
}

type myinfo struct {
	//ml sync.Mutex
	info []byte
}

func (this *myinfo) init()  {
	this.info = make([]byte,0)
}

func (this *myinfo) addInfo(val byte)  {
	//this.ml.Lock()
	//defer this.ml.Unlock()
	this.info = append(this.info,val)
}

func (this *myinfo) print() {
	fmt.Println(this.info)
}


var (
	in1 chan int
	in2 chan int
)

func info1(info1 *myinfo ) {
	info1.addInfo('x')
	in1 <- 1
}
func info2(info2 *myinfo) {
	info2.addInfo('y')
	in2 <- 1
}
func main() {
	in1 = make(chan int)
	in2 = make(chan int)

	var mytestInfo myinfo
	mytestInfo.init()

	go info1(&mytestInfo)
	go info2(&mytestInfo)

	select {
	case <-in1:
		fmt.Println("the info1 xxxx")
	}

	select {
	case <-in2:
		fmt.Println("the info2 yyyy")
	}

	mytestInfo.print()


	//ll.Lock()
	go f()
	//ll.Lock()
	fmt.Println(aa)
	//ll.Unlock()

	//var input string
	//fmt.Scanln(&input)	//这里的 Scanln 代码需要我们在程序退出前按下任意键结束。
	fmt.Println("done")

}