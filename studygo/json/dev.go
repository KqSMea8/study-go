package main

import(
	"encoding/json"
	//"strings"
	"fmt"
	"os"
)

type person struct{
	name string  `json:"name"`
	age int  `json:"age"`
	weight float32  `json:"wieght"`
}

func main() {
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	var students [](map[string]person)
	for i:=0;i<10;i++ {
		pstring := fmt.Sprintf("abc%d",i)
		tmp:= make(map[string]person)
		var xtmp person
		//println(pstring)
		xtmp = person{"xiao", 22, 58.60}
		tmp[pstring] = xtmp
		students = append(students,tmp)
	}

	for spk,record := range students {
		pk := fmt.Sprintf("abc%d",spk)
		//println(spk,pk)
		fmt.Println(record[pk].name, record[pk].age,record[pk].weight)
		for _, pvs := range record {
			fmt.Println(pvs.name)
		}
	}

	rb := json.NewEncoder(os.Stdout);
	rb.Encode(students)
	//fmt.Println(string(rb))
}

