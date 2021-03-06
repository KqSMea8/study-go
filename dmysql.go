package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

//显示多条数据
type myperson struct {
	Id uint64 `sql:"id"`
	Age uint8 `sql:"age"`
	Name string `sql:"name"`
}

func main() {

	//链接数据库
	db,err := sql.Open("mysql","xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//插入数据
	var result sql.Result
	result,err = db.Exec("insert into infos (age,name) values (?,?)",27,"肖建和")
	if err !=nil {
		fmt.Println(err)
		return
	}
	lastId,_:= result.LastInsertId()
	fmt.Println("新插入记录id为",lastId)

	//显示单条数据
	fmt.Println("获取单条记录")
	var row *sql.Row
	row = db.QueryRow("select * from infos")
	var id,age uint64
	var name string
	err = row.Scan(&id,&name,&age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id,"\t",name,"\t",age)

	fmt.Println("获取多条记录")
	var rows *sql.Rows
	rows,err = db.Query("select id,age,name from infos")
	if err != nil {
		fmt.Println(err)
	}

	myp:=myperson{}

	testStructObject(rows,myp)

	//fmt.Printf("这是一个Mysql例子\n",ssvalues)
}

func testStructObject(rows *sql.Rows,myp interface{}) (mytest []interface{}) {
	columns,_:=rows.Columns()
	svalues :=make([]interface{},len(columns))
	reStruct:=reflect.ValueOf(&myp)
	fmt.Println("what:",columns)
	for si,sv:=range columns{
		//psv := reStruct.Type().Field(si).Tag.Get("sql")
		//fmt.Println("hello1:",sv,psv)
		pname :=findTagName(myp,sv)
		//if pname==""{
		//	continue
		//}
		fmt.Println("pname:",pname)
		//return
		svalues[si] = reStruct.Elem().Interface()
		fmt.Println("ppp:",svalues)
		//svalues[si] = reStruct.Elem().FieldByName(pname).Interface()
		//svalues[si] = reStruct.Field(si).Addr().Interface()
		//fmt.Println("hello:",reStruct.Elem().Field(si).Addr().Interface())
	}
	//return
	//mytest:=[]myperson
	for rows.Next() {
		rows.Scan(svalues...)
		//results[ii] = svalues
		//ii++
		mytest = append(mytest,myp)
		//fmt.Println(myp.Age,myp.Id,myp.Name)
	}
	//fmt.Println(mytest[0].Age)

	for ppi:=0;ppi<len(mytest);ppi++ {
		fmt.Println(mytest[ppi])
		//fmt.Println(mytest[ppi].Id,mytest[ppi].Age,mytest[ppi].Name)
	}
	return
}

func findTagName(myp interface{},tag string) (name string) {
	myt:=reflect.TypeOf(myp)
	//myv:=reflect.ValueOf(&myp)

	for mykk:=0;mykk<myt.NumField();mykk++ {
		ptag:=myt.Field(mykk).Tag.Get("sql")
		if tag == ptag {
			name = myt.Field(mykk).Name
			//vale = myv.Elem()
			break
		}
		//fmt.Printf("%s -- %v --tag %v \n", myt.Field(mykk).Name, myv.Field(mykk).Interface(),ptag)
	}
	return
}


/*
mygo 下需要的表
CREATE TABLE `my` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `age` int(10) unsigned NOT NULL DEFAULT '0',
  `name` varchar(45) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 |
 */