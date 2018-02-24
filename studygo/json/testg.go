package main

import(
	"encoding/json"
	//"strings"
	"fmt"
	//"os"
)
/**
带key的输出
 */
type Person struct{
	Name string  `json:"name" db:"name"`
	Age int  `json:"age_id" db:"age"`
	Weight float32  `json:"wieght_id"  db:"weight"`
}

func main() {
	slcD := &Person{}
	slcB, _ := json.Marshal(slcD)
	var pb map[string]interface{}
	_ = json.Unmarshal(slcB,&pb)
	var args []interface{}
	var keys []string
	for key, val := range pb {
		keys = append(keys,key)
		args = append(args, val)
	}
	fmt.Println(pb)
	fmt.Println(keys)
	fmt.Println(args)
}

