package main

import (
	"github.com/huandu/go-sqlbuilder"
	"fmt"
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MessageDao struct {
	ID         int64  `thrift:"id,1,required" db:"id" json:"id"`
	RiderId    int64  `thrift:"riderId,2,required" db:"riderId" json:"riderId"`
	MasterId   int64  `thrift:"masterId,3" db:"masterId" json:"masterId,omitempty"`
	Title      string `thrift:"title,4,required" db:"title" json:"title"`
	SubTitle   string `thrift:"subTitle,5" db:"subTitle" json:"subTitle,omitempty"`
	Content    string `thrift:"content,6" db:"content" json:"content,omitempty"`
	Type       int64  `thrift:"type,7" db:"type" json:"type,omitempty"`
	CreateTime int64  `thrift:"createTime,8" db:"createTime" json:"createTime,omitempty"`
	UpdateTime int64  `thrift:"updateTime,9" db:"updateTime" json:"updateTime,omitempty"`
	Status     int64  `thrift:"status,10" db:"status" json:"status,omitempty"`
	Tts        string `thrift:"tts,11" db:"tts" json:"tts,omitempty"`
	JumpType   string `thrift:"jumpType,12" db:"jumpType" json:"jumpType,omitempty"`
	JumpUrl    string `thrift:"jumpUrl,13" db:"jumpUrl" json:"jumpUrl,omitempty"`
	IsDel      int64  `thrift:"isDel,14" db:"isDel" json:"isDel,omitempty"`
	Phone      string `thrift:"phone,15" db:"phone" json:"phone,omitempty"`
	PicKey     string `thrift:"picKey,16" db:"picKey" json:"picKey,omitempty"`
	IconType   string `thrift:"iconType,17" db:"iconType" json:"iconType,omitempty"`
	ExtId      int64  `thrift:"extId,18" db:"extId" json:"extId,omitempty"`
}

//实例化
func NewMesageDao() *MessageDao {
	return new(MessageDao)
}

//定义表
func (md *MessageDao) GetTableName() string {
	return "soda_message_rider"
}

//添加记录
func (md *MessageDao) Insert() (re int64, err error) {
	return re,err
}

//修改内容
func (md *MessageDao) Update() (re int64, err error) {
	return re,err
}

//获取单条
func (md *MessageDao) SelectOne(ctx context.Context,col []string,andExpr map[string]interface{}) (re []*MessageDao, err error) {

	sqlHandle := sqlbuilder.NewSelectBuilder()
	sqlHandle.Select(col ...)
	sqlHandle.From(md.GetTableName())
	for skey, sval := range andExpr {
		sqlHandle.Where(sqlHandle.Equal(skey,sval))
	}
	sqlHandle.OrderBy("id").Desc()
	sql, args := sqlHandle.Build()
	//sql = sql + " ORDER BY `id` DESC"
	fmt.Println(sql)
	fmt.Println(args)

	return re,err
}

//获取多条
func (md *MessageDao) SelectMulti(ctx context.Context,col []string,andExpr []string) (re []*MessageDao, err error) {
	sqlHandle := sqlbuilder.NewSelectBuilder()

	sqlHandle.Select(col ...)
	sqlHandle.From(md.GetTableName())
	sqlHandle.Where(andExpr ...)


	//return re,err
	return re,err
}

//按条件删除
func (md *MessageDao) DeleteConModi() (re int64,err error) {
	return re,err
}

func main() {

	//链接数据库
	db,err := sql.Open("mysql","xxx:xxxx@tcp(xx.95.136.xxxx:3306)/mygo?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	ctx := context.Background()
	test:= &MessageDao{}
	col := []string{
		"id","rider_id",
	}
	andpx := map[string]interface{} {
		"id":1,
		"roder_id":3,
	}

	test.SelectOne(ctx,col,andpx)
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "name", "COUNT(*) as c")
	sb.From("user")
	sb.Where(sb.Equal("id",2))
	sb.Where(sb.G("rider_id",2))//,sb.E("master_id",2),sb.In("status", 1, 2, 5))
	sb.Asc()

	sql, args := sb.Build()
	fmt.Println(sql)
	fmt.Println(args)
}