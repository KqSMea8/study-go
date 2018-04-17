package main

/**
  myint 和 int 之间转化的例子

*/
import(
	//"unsafe"
	"fmt"
	"unsafe"
)

func main() {
	type myInt int

	a := []myInt{0,1,2}
	//b := a
	b := *(*[]int)(unsafe.Pointer(&a))
	fmt.Println(b)

	b[0] = 3

	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

}
