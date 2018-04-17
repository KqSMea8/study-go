package main

import(
	"time"
	"math/rand"
	"fmt"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println("/")
	fmt.Println(time.Now().UnixNano())
	fmt.Printf("%x%X", rand.Int63(), rand.Int63())
	fmt.Println("/")
	fmt.Println(rand.Int63())
	fmt.Printf("%.3s","我们的人呀")
}

