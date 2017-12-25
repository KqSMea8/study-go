package main

import "fmt"

func main() {
	Minimum(10, 15, 32, 46, 2, 3)  //1
	var sliceInt = []int{10, 15, 32, 46, 2, 3}
	sliceInt2 := make([]interface{},0)
	sliceInt2 = append(sliceInt2,10, 15, 32, 46, 2, 1)
	Minimum(sliceInt[0], sliceInt[1], sliceInt[2], sliceInt[3], sliceInt[4], sliceInt[5])  //2
	Minimum(sliceInt2[0],sliceInt2[1:]... )  //3
}

func Minimum(first interface{}, rest ...interface{})  {
	min := first
	for _, value := range rest {
		switch value := value.(type) {
		case int:
			if value < min.(int) {
				min = value
			}
		case float32:
			if value < min.(float32) {
				min = value
			}
		case float64:
			if value < min.(float64) {
				min = value
			}
		case string:
			if value < min.(string) {
				min = value
			}
		}
	}
	fmt.Println(min)
}
