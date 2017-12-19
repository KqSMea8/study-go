package main

import(
	"fmt"
	"html/template"
	"log"
	"os"
)

//设定模版
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt}} days
{{end}}`

type User struct {
	Number int
	User string
	Title string
	CreatedAt string
}

func main() {
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		TotalCount int
		Items map[int]*User //用指针类型
	}

	tmp := map[int]*User{}
	tmp[0] = &User{100,"xiaojh","test","31"}
	tmp[1] = &User{101,"xiaojh2","test2","32"}

	fmt.Sprintf(tmp)
	data.TotalCount = 2
	data.Items = tmp
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
	fmt.Println("--------ok--------")
	//fmt.Println(templ)
}