package main

import (
	"fmt"
)

func main() {

	in1 := make(chan int)
	in2 := make(chan int)
	in3 := make(chan int)
	done:= make(chan bool)
	var pp,cc,dd int
	go func(msg string) {
		for i:=0;i<3;i++ {
			if i==0 {
				fmt.Println(i)
				pp++
				in1 <- 0
			} else if i==1 {
				fmt.Println(i)
				pp++
				in2 <- 1
			} else {
				fmt.Println(i)
				dd++
				in3 <- 2
			}

			fmt.Println(msg,i)  //注意这个位置
		}
		done <- true
	}("going")

	for {
		fmt.Println("out",in1,in2,in3)

		select {
		case <-in1:
			fmt.Println("in0 is ok pp is ",pp)
		case <-in2:
			fmt.Println("in1 is ok pp and cc is ",pp,cc)
		case <-in3:
			fmt.Println("in2 is ok dd is ",dd)
		case <-done:
			return
		}
	}

	var input string
	fmt.Scanln(&input)	//这里的 Scanln 代码需要我们在程序退出前按下任意键结束。
	fmt.Println("done")

}