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

	var students JsonPerson
	for i:=0;i<10;i++ {
		//pstring := fmt.Sprintf("abc%d",i)
		//tmp:= make(map[string]person)
		//var xtmp person
		//println(pstring)
		//xtmp = person{"xiao", 22, 58.60}
		tmp := Person{Name:"xiao",Age: 22, Weight:58.60}
		students.Infos = append(students.Infos,tmp)
	}

	for _,record := range students.Infos {
		fmt.Println(record)
		//pk := fmt.Sprintf("abc%d",spk)
		//println(spk,pk)
		//fmt.Println(record[pk].name, record[pk].age,record[pk].weight)
		fmt.Println(record.Name, record.Age,record.Weight)
		/*for _, pvs := range record {
			fmt.Println(pvs)
		}*/
	}

	//rb := json.NewEncoder(os.Stdout);
	//rb.Encode(students.infos)

	b, err := json.Marshal(students)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//fmt.Println(string(rb))


	type Server struct {
		ServerName string
		ServerIP   string
	}

	type Serverslice struct {
		Servers []Server
	}

	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	bx, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(bx))

}

