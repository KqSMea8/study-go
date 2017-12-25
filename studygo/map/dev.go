package main

import(

)
import (
	"sort"
	"fmt"
)

func main() {
	var m map[string]int
	m = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16,
		"indio": 87, "juliet": 65, "kilo": 43, "lima": 98,
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, sk := range keys {
		v := m[sk]
		fmt.Println("key:%s=>%d",sk,v)
	}
}

