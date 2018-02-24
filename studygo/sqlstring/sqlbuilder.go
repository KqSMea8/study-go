// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

func main() {
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "name", "COUNT(*) as c")
	sb.From("user")
	sb.Where(sb.In("status", 1, 2, 5))

	sql, args := sb.Build()
	fmt.Println(sql)
	fmt.Println(args)
}
