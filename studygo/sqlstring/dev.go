// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import (
	"encoding/json"
	"strings"
	"fmt"
	"regexp"
)

type Response struct {
	RiderID int    `json:riderId`
	UserId  int    `json:userId`
	Name    string `json:"name"`
	Age     int    `json:"age"`
}

type Point struct {
	ID         int64   `thrift:"id,1,required" db:"id" json:"id"`
	RiderId    int64   `thrift:"riderId,2,required" db:"riderId" json:"riderId"`
	Name       string  `thrift:"name,3" db:"name" json:"name,omitempty"`
	Lat        float64 `thrift:"lat,4,required" db:"lat" json:"lat"`
	Lng        float64 `thrift:"lng,5,required" db:"lng" json:"lng"`
	PoiNo      int64   `thrift:"poiNo,6" db:"poiNo" json:"poiNo,omitempty"`
	PoiName    string  `thrift:"poiName,7" db:"poiName" json:"poiName,omitempty"`
	PoiRange   int32   `thrift:"poiRange,8,required" db:"poiRange" json:"poiRange"`
	CityID     int32   `thrift:"cityID,9" db:"cityID" json:"cityID,omitempty"`
	CityName   string  `thrift:"cityName,10" db:"cityName" json:"cityName,omitempty"`
	CreateTime int64   `thrift:"createTime,11" db:"createTime" json:"createTime,omitempty"`
	UpdateTime int64   `thrift:"updateTime,12" db:"updateTime" json:"updateTime,omitempty"`
}

func Insert(tableName string, obj interface{}) (sql string, args []interface{}) {
	var placeholders []string
	var keys []string
	var params map[string]interface{}

	js, _ := json.Marshal(obj)
	json.Unmarshal(js, &params)

	for key, val := range params {
		keys = append(keys, underscore(key))

		args = append(args, val)
		placeholders = append(placeholders, "?")
	}
	sql = fmt.Sprintf("INSERT INTO `%v` (`%v`) VALUES (%v)",
		tableName,
		strings.Join(keys, "`, `"),
		strings.Join(placeholders, ", "))

	return sql, args
}
func CamelCase(in string) string {
	tokens := strings.Split(in, "_")
	for i := range tokens {
		tokens[i] = strings.Title(strings.Trim(tokens[i], " "))
	}
	return strings.Join(tokens, "")
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func underscore(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
func main() {
	obj := Response{
		RiderID: 1,
		UserId:  2,
		Name:    "apple",
	}

	insertSql, args := Insert("xxx", obj)
	fmt.Println(insertSql)
	fmt.Println(args)
	fmt.Println("____")

	point := Point{
		ID:         111,
		RiderId:    222,
		Name:       "point",
		Lat:        12.34,
		Lng:        45.67,
		PoiName:    "PointName",
		PoiRange:   100,
		CityID:     1,
		CityName:   "city",
		CreateTime: 4321,
		UpdateTime: 12345,
	}

	insertSql, args = Insert("xxx", point)
	fmt.Println(insertSql)
	fmt.Println(args)
	fmt.Println("____")

	// row scan出来的结果, 不要往后面追加一堆 &obj.field1, &obj.field2
	obj1 := Response{}
	params := map[string]interface{}{"user_id": 2, "rider_id": 22, "name": "banana", "age": 10}
	m := make(map[string]interface{})
	for k, v := range params {
		key := CamelCase(k)
		m[key] = v
	}
	js, _ := json.Marshal(m)
	json.Unmarshal([]byte(js), &obj1)
	fmt.Printf("%#v", obj1)
}
