package main

import (
	"bufio"
	"os"
	"fmt"
)
func main() {
	var s,sep string
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for i:=0;i<len(counts)-1;counts{

	}
	fmt.Println(s)

}
