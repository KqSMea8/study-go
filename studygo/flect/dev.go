package main

import(
	"reflect"
	"fmt"
)

type Foo struct {
	X string
	Y string
}
func main() {
	var i int
	i = 123
	var f Foo
	f = Foo{"abc","def"}
	xm :=map[string]string{"a":"xiaojh","b":"china"}

	println(reflect.TypeOf(i).Name(),reflect.TypeOf(f).Name(),reflect.TypeOf(i).Kind().String())
	xf :=reflect.ValueOf(&f).Elem()  //必须用取址
	fmt.Println(xf)
	fmt.Println(reflect.ValueOf(i).String(),reflect.ValueOf(f).String(),reflect.ValueOf(&xm).Elem())
}
