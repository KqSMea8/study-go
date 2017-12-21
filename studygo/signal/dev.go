package main

import(
	"fmt"
	"os"
	"os/signal"
)

func main() {
	test := make(chan os.Signal)
	signal.Notify(test)
	fmt.Println("启动")
	out := <-test
	fmt.Println("退出信息:")
}

