package main

import(
	"encoding/json"
	//"strings"
	"fmt"
	"os"
)
/**
带key的输出
 */
type Person struct{
	Name string  `json:"name"`
	Age int  `json:"age"`
	Weight float32  `json:"wieght"`
}

func main() {
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	var students [](map[string]Person)
	for i:=0;i<10;i++ {
		pstring := fmt.Sprintf("abc%d",i)
		tmp:= make(map[string]Person)
		var xtmp Person
		//println(pstring)
		xtmp = Person{"xiao", 22, 58.60}
		tmp[pstring] = xtmp
		students = append(students,tmp)
	}

	for spk,record := range students {
		pk := fmt.Sprintf("abc%d",spk)
		fmt.Println(record[pk].Name, record[pk].Age,record[pk].Weight)
		for _, pvs := range record {
			fmt.Println(pvs.Name)
		}
	}

	//一种
	rb := json.NewEncoder(os.Stdout);
	rb.Encode(students)
	//fmt.Println(string(rb))

	//二种
	b, err := json.Marshal(students)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

