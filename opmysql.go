package main

import (
	"context"
	"fmt"
	"time"

)

// 模拟一个最小执行时间的阻塞函数
func inc(a int) int {
	res := a + 1                // 虽然我只做了一次简单的 +1 的运算,
	time.Sleep(1 * time.Second) // 但是由于我的机器指令集中没有这条指令,
	// 所以在我执行了 1000000000 条机器指令, 续了 1s 之后, 我才终于得到结果。B)
	return res
}

// 向外部提供的阻塞接口
// 计算 a + b, 注意 a, b 均不能为负
// 如果计算被中断, 则返回 -1
func Add(ctx context.Context, a, b int) int {
	res := 0
	for i := 0; i < a; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}
	for i := 0; i < b; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}
	return res
}

// MyError is an error implementation that includes a time and message.
type TimeoutError struct {
	When time.Time
}

func (e TimeoutError) Error() string {
	return fmt.Sprintf("The timeout occur at : %v ", e.When)
}

func toerror() error {
	return TimeoutError{
		time.Now(),
	}
}
func Example() {
	if _, ok := toerror().(TimeoutError); ok {
		println(toerror().Error())
	}
	// Output: 1989-03-15 22:30:00 +0000 UTC: the file system has gone away
}

const ConfFile = "./config.ini"

type Config struct {
	Logconf   log.Config   `toml:"log"`
	Mysqlconf mysql.Config `toml:"mysql"`
}

//事务会在commit方法中自动释放链接资源，无需手动关闭，也没有手动关闭的方法。
func testTx(factory *mysql.Factory, ctx context.Context) {
	tx, err := factory.New(ctx).Begin(nil)
	if err != nil {
		println("TX:", err.Error())
	}
	if tx != nil {
		println("Begin TX")
		rs, exerr := tx.Exec("INSERT INTO t_table (`desc`)  VALUES('tx test')")
		if exerr != nil {
			println("Begin:", exerr.Error())
		}
		if rs != nil {
		}
		tx.Commit()
	}
}

//事务会在Rollback方法中自动释放链接资源，无需手动关闭，也没有手动关闭的方法。
func testTxRB(factory *mysql.Factory, ctx context.Context) {
	tx, err := factory.New(ctx).Begin(nil)
	if err != nil {
		println("TX:", err.Error())
	}
	if tx != nil {
		println("Begin TX")
		rs, exerr := tx.Exec("INSERT INTO t_table (`desc`)  VALUES('tx test')")
		if exerr != nil {
			println("Begin:", exerr.Error())
		}
		if rs != nil {
		}
		tx.Rollback()
	}
}

// 不会自动释放资源，必须手动close()释放链接
// defer stmt.Close()
// 如果需要执行多条sql语句，这些SQL的结构相同只是where条件的参数不同
// 或者只是insert/update的值不同,这种情况下，推荐使用Prepare进行操作，效率更高。
func tesTxPrepare(factory *mysql.Factory, ctx context.Context) {
	tx, err := factory.New(ctx).Begin(nil)
	stmt, err := tx.Prepare("INSERT INTO t_table(`desc`) VALUE(?)")
	defer stmt.Close()
	if err != nil {
		println("testPrepare:" + err.Error())
	} else {
		res, err := stmt.Exec("tesTxPrepare")

		if err != nil {
			println("stmt err:", err.Error())
		}
		if res != nil {
			var num int64
			num, _ = res.RowsAffected()
			println("stmt RowsAffected:", num)
		}

		res, err = stmt.Exec("tesTxPrepare")
		if err != nil {
			println("stmt err:", err.Error())
		}
		if res != nil {
			var num int64
			num, _ = res.RowsAffected()
			println("stmt RowsAffected:", num)
		}
		tx.Commit()
	}
}

// Query不会自动释放资源，必须手动close()释放链接
// defer rows.Close()
func testTxQuery(factory *mysql.Factory, ctx context.Context) {
	tx, err := factory.New(ctx).Begin(nil)
	rows, err := tx.Query("select * from t_table")
	defer rows.Close()
	if err != nil {
		println("query:" + err.Error())
		return
	}
	for {
		ok, err := rows.Next()
		if err != nil {
			println("query:" + err.Error())
			break
		}
		if !ok {
			break
		}
		var setName, localName string
		rows.Scan(&setName, &localName)
		println(setName, localName)
	}
	tx.Commit()
}

// 当你确定当前的查询语句最多返回一条语句的情况下，使用本函数
// QueryRow会自动释放资源,无需手动处理，也没有提供相应的方法
func testTxQueryRow(factory *mysql.Factory, ctx context.Context) {
	tx, err := factory.New(ctx).Begin(nil)
	row, err := tx.QueryRow("select * from t_table where id=1")
	if err != nil {
		println("testQueryRow:" + err.Error())
		return
	}
	var setName, localName string
	err = row.Scan(&setName, &localName)
	println(setName, localName)
	if err != nil {
		println("testQueryRow:" + err.Error())
		return
	}
	tx.Commit()
}

// Exec会在方法中自动释放链接资源，无需手动关闭，也没有手动关闭的方法。
func testExec(factory *mysql.Factory, ctx context.Context) {
	res, err := factory.New(ctx).Exec("INSERT INTO t_table VALUES( ？)", "test Exec")
	if err != nil {
		println("RowsAffected:", err.Error())
	}
	if res != nil {
		var num int64
		num, _ = res.RowsAffected()
		println("RowsAffected:", num)
	}
}

// Query不会自动释放资源，必须手动close()释放链接
// defer rows.Close()
func testQuery(factory *mysql.Factory, ctx context.Context) {
	rows, err := factory.New(ctx).Query("select * from t_table")
	defer rows.Close()
	if err != nil {
		println("query:" + err.Error())
		return
	}
	for {
		has, err := rows.Next()
		if err != nil {
			println("query:" + err.Error())
			return
		}
		if !has {
			break
		}

		var setName, localName string
		rows.Scan(&setName, &localName)
		println(setName, localName)
	}
}

// 当你确定当前的查询语句最多返回一条语句的情况下，使用本函数
// QueryRow会自动释放资源,无需手动处理，也没有提供相应的方法
func testQueryRow(factory *mysql.Factory, ctx context.Context) {
	row, err := factory.New(ctx).QueryRow("select * from t_table where id=1")
	if err != nil {
		println("testQueryRow:" + err.Error())
		return
	}
	var setName, localName string
	err = row.Scan(&setName, &localName)
	println(setName, localName)
	if err != nil {
		println("testQueryRow:" + err.Error())
		return
	}
}

// 链接池的全局关闭方法，请不要轻易调用，除非你真的知道调用的理由
func testClose(factory *mysql.Factory, ctx context.Context) {
	err := factory.New(ctx).Close()
	if err != nil {
		println("Close:" + err.Error())
	}
}

func testPing(factory *mysql.Factory, ctx context.Context) {
	err := factory.New(ctx).Ping()
	if err != nil {
		println("Ping:" + err.Error())
	} else {
		println("the db connection is ok")
	}
}

// 不会自动释放资源，必须手动close()释放链接
// defer stmt.Close()
// 如果需要执行多条sql语句，这些SQL的结构相同只是where条件的参数不同
// 或者只是insert/update的值不同,这种情况下，推荐使用Prepare进行操作，效率更高。
func testPrepare(factory *mysql.Factory, ctx context.Context) {
	stmt, err := factory.New(ctx).Prepare("INSERT INTO t_table(`desc`) VALUE(?)")
	defer stmt.Close()
	if err != nil {
		println("testPrepare:" + err.Error())
	} else {
		res, err := stmt.Exec("Prepare")

		if err != nil {
			println("stmt err:", err.Error())
		}
		if res != nil {
			var num int64
			num, _ = res.RowsAffected()
			println("stmt RowsAffected:", num)
		}

		res, err = stmt.Exec("Prepare")
		if err != nil {
			println("stmt err:", err.Error())
		}
		if res != nil {
			var num int64
			num, _ = res.RowsAffected()
			println("stmt RowsAffected:", num)
		}
	}
}

func testStats(factory *mysql.Factory, ctx context.Context) {
	stat, err := factory.New(ctx).Stats()
	if err != nil {
		println("testPrepare:" + err.Error())
	} else {
		println("OpenConnections:", stat.OpenConnections)
	}
}

var config Config
var factory *mysql.Factory

func init() {
	// init the config, only need init once in any project
	if _, err := toml.DecodeFile(ConfFile, &config); err != nil {
		fmt.Printf("fail to read config.||err=%v||config=%v", err, ConfFile)
		return
	}
	// init the sql Factory, only need init once in any project
	log.Init(&config.Logconf)
	factory = mysql.NewFactory()
	if err := factory.Open(&config.Mysqlconf); err != nil {
		fmt.Printf("fail to get sql Factory instantce || err=%v", err)
		factory = nil
	}

}

func main() {

	defer log.Close()

	log.Infof("begin!")

	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	testPing(factory, ctx)
	//testTx(factory,ctx)
	//testQueryRow(factory,ctx)
	//testTxRB(factory,ctx)
	tesTxPrepare(factory, ctx)
	//testPrepare(factory,ctx)
	testQuery(factory, ctx)
	//testTxQuery(factory,ctx)
	//testTxQueryRow(factory,ctx)
	//testClose(factory,ctx)
	//testPing(factory,ctx)
	testStats(factory, ctx)
	//go testQueryRow(factory,ctx)

	//_,err = res.RowsAffected()
	//println("RowsAffected:",err)

	retry := 0
	if retry < 1 {
		retry = 1
	}

	println("retry:", retry)

	Example()

	LABEL:
		select {
		case <-time.After(1000000000 * time.Nanosecond):
			println("test goto and time after")
			goto LABEL
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

	{
		// 使用开放的 API 计算 a+b
		a := 1
		b := 2
		timeout := 2 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		res := Add(ctx, 1, 2)
		fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
	}
	{
		// 手动取消
		a := 1
		b := 2
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			cancel() // 在调用处主动取消
		}()
		res := Add(ctx, 1, 2)
		fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
	}
}
