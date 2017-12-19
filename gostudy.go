package main

import (
	//"fmt"
	"unicode/utf8"
)

type e interface {

}

func main() {

	var a=[...]int{0,1,2,3,4,5,6,7}
	var s = make([]int, 2)
	copy(s, a[0:])
	n2 := copy(s, s[2:])
	for _,vs := range  s {
		println(vs)
	}

	for i:=1;i<2;i++ {
		println(i)
	}
	mystring := "xpfadg,agagag"
	ab := utf8.RuneCount([]byte(mystring))
	println(ab);
	println(n2)
	myfunc()
}

func myfunc() {
	i := 0
	Here:
	println(i)
	i++
	if i<10 {
		goto Here
	}
	//goto Here
	//Stop:
}

func averageFloat(xf []float64) (avg float64) {
	sum := 0.0
	howlen :=len(xf)
	switch howlen {
	case 0:
		avg=0
	default:
		for _,v:=range xf {
			sum+=v
		}
		avg = sum/float64(howlen)
	}
	return avg
}

func renab(a int ,b int)(int,int) {
	if (a>b) {
		return b,a
	}
	return a,b
}

func mult2(f e) e {
	switch f.(type) {
	case int:
		return f.(int)*2
	case string:
		return f.(string)+f.(string)+f.(string)+f.(string)
	}
	return f
}
