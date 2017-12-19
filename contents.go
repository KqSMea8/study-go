package main

import(
	"context"
	"log"
	"os"
	"time"
)

var slog *log.Logger
func someHandler() {
	ctx,cancle := context.WithCancel(context.Background())
	go doStuff(ctx)
	//10秒钟后取消doStuff
	time.Sleep(10*time.Second)
	cancle()
}
func timeoutHandlerWithDead() {
	ctx,cancle := context.WithDeadline(context.Background(),time.Now().Add(5*time.Second))
	go doStuff(ctx)
	time.Sleep(10*time.Second)
	cancle()
}
func timeoutHandlerWithtime() {
	ctx,cancle := context.WithTimeout(context.Background(),5*time.Second)
	go doTimeOutStuff(ctx)
	time.Sleep(10*time.Second)
	cancle()
}
func doStuff(ctx context.Context)  {
	for {
		time.Sleep(1*time.Second)
		select {
		case <- ctx.Done():
			slog.Printf("done")
			return
		default:
			slog.Printf("work")
		}
	}
}
func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1*time.Second)
		if deadline,ok:=ctx.Deadline();ok {
			slog.Printf("deadline set")
			if time.Now().After(deadline) {
				slog.Printf(ctx.Err().Error())
				return
			}
		}
		select {
		case <- ctx.Done():
			slog.Printf("done")
			return
		default:
			slog.Printf("work")
		}
	}
}
func main() {
	slog = log.New(os.Stdin,"",log.Ltime)
	slog.Printf("第一类 withcancle")
	someHandler()  //第一种
	slog.Printf("第二类 withdeadline")
	//timeoutHandlerWithDead()  //第二种
	slog.Printf("第三类 withtimeout")
	//timeoutHandlerWithtime()   //第三种
	slog.Printf("down")
}
