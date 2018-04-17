package xtest

import (
	"testing"
	"time"
)


//线上压力测试
func Benchmark_Pushtest(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		time.Sleep(1*time.Second) //线下是1秒5次.别太快
		//XPush()
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	// ...

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		time.Sleep(1*time.Second)
		//XPush()
	}
}

