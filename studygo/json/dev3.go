package main

import(
	"encoding/json"
	//"strings"
	"fmt"
	//"os"
)
/**
数值自增的 key输出
 */
type Person struct{
	Name string  `json:"Name"`
	Age int  `json:"Age"`
	Weight float32  `json:"Weight"`
}

type JsonPerson struct {
	Infos []Person
}

func main() {
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	slccc := "nice to see you"
	slaaa, _ := json.Marshal(slccc)
	scla := []byte(slccc)
	abb:=string(scla)
	acc:=string(slcB)
	fmt.Println(abb,acc,string(slaaa))

	/*var jjj = []byte (`  [
        { "Space" : "YCbCr" , "Point" : { "Y" : 255 , "Cb" : 0 , "Cr" : - 10 } } ,
        { "Space" : "RGB" ,   "Point" : { "R" : 98 , "G" : 218 , "B" : 255 } }
    ] ` )
	fmt.Println(string(jjj))*/


}

