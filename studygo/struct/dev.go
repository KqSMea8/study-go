package main

import(
)
import "fmt"

type User struct {
	name string
	pass string
}

func main() {
	mytest:=User{name:"xiaojh",pass:"1"}
	convStruct(&mytest)
}
func convStruct(mu interface{}) {
	vas:=mu.(*User)  //显未转
	fmt.Printf("===%v__%v\n",mu,vas.pass)
}
