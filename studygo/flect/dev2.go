package main

import (
	"fmt"
	"reflect"
)

type User struct  {
	Id int `thrift:"id,1,required" db:"id" json:"id"`
	Name string `thrift:"Name,1,required" db:"Name" json:"id"`
	Addr string `thrift:"Addr,1,required" db:"Addr" json:"id"`
}

func main(){
	u := User{ Name:"xiaojh", Addr:"china",Id:1001}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	colVal := []interface{}{}
	colVal = append(colVal,int(100))
	for k := 1; k < t.NumField(); k++ {
		//fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
		colVal = append(colVal, v.Field(k).Interface())
	}
	var xu *User
	xu = &User{ Name:"xiaojh", Addr:"china",Id:1001}
	xt := reflect.TypeOf(xu).Elem()
	xv := reflect.ValueOf(xu).Elem()
	xcolVal := []interface{}{}
	xcolVal = append(xcolVal,int(100))
	for xk := 1; xk < xt.NumField(); xk++ {
		fmt.Printf("%s -- %v \n", xt.Field(xk).Name, v.Field(xk).Interface())
		xcolVal = append(xcolVal, xv.Field(xk).Interface())
	}
	fmt.Printf("%v --\n",xcolVal)
} 