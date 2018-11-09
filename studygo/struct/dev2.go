package main

import(
)
import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	pass string
}

func main() {
	mytest:=User{name:"xiaojh",pass:"100"}
	convStruct(mytest)
}
func convStruct(mu interface{}) {
	smu :=[]string{
		"name","pass",
	}
	values := reflect.ValueOf(&mu)//.Interface()
	value := values.Elem()
	for sdv,spv:=range smu {
		//tname := "name"
		pvs := value.FieldByName(spv)
		fmt.Printf("===%v__%v_%v__%v__%v\n",sdv,spv, values, value, pvs)
	}
}
