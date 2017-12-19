package main

import(
	"fmt"
	"flag"
	//"log"
)

func main() {
	data_path := flag.String("name","sample","DB data path")
	log_file := flag.String("l","long file","log file")
	nowait_flag := flag.Bool("W",false,"if write to into")
	flag.Parse()

	fmt.Printf("name:",*data_path)
	fmt.Printf("l:",*log_file)
	fmt.Printf("w",*nowait_flag)
}
