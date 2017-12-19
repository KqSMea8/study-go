package main

import(
	"./stack"
)

func main() {
	//xdata :=[10]int{1,2,3,4,5,6,7,8,9,10}
	//bstatck  := new(stack.Stack)
	bstatck := &stack.Stack{}
	bstatck.Xpush(1)
	bstatck.Xpush(2)
	bstatck.Xpush(4)

	_,data := bstatck.Xget()

	xgetval(data)

	a:=bstatck.Xpop()
	b:=bstatck.Xpop()
	//println(a)
	println("splite\n",a,b)

	_,sdata := bstatck.Xget()
	xgetval(sdata)
}

func xgetval(sdata [10]int) {
	for _,pv:=range sdata {
		if pv==0 {
			//continue
		}
		println(pv)
	}
}