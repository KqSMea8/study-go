package main
import (
	//"database/sql"
	"fmt"
	"github.com/cznic/ql"
	_ "github.com/go-sql-driver/mysql"
)

//显示多条数据
type myperson struct {
	Id uint64 `sql:"id"`
	Age uint8 `sql:"age"`
	Name string `sql:"name"`
}

func main() {

	//链接数据库
	/*db,err := sql.Open("mysql","xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close() */

	qldb, err := ql.OpenMem()
	//ql.OpenFile()
	if err != nil {
		panic(err)
	}

	rss, _, err := qldb.Run(ql.NewRWCtx(), `
	BEGIN TRANSACTION;
		CREATE TABLE foo (i int);
		INSERT INTO foo VALUES (10), (20);
		CREATE TABLE bar (fooID int, s string);
		INSERT INTO bar SELECT id(), "ten" FROM foo WHERE i == 10;
		INSERT INTO bar SELECT id(), "twenty" FROM foo WHERE i == 20;
	COMMIT;
	SELECT *
	FROM foo, bar
	WHERE bar.fooID == id(foo)
	ORDER BY id(foo);`,
	)
	if err != nil {
		panic(err)
	}

	for _, rs := range rss {
		if err := rs.Do(false, func(data []interface{}) (bool, error) {
			fmt.Println(data)
			return true, nil
		}); err != nil {
			panic(err)
		}
		fmt.Println("----")
	}
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