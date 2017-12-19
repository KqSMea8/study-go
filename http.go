package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		resp,error := http.Get(url)
		if error != nil {
			fmt.Fprint(os.Stderr,"fetch:%v \n",error)
			os.Exit(1)
		}
		sresult,error := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if error != nil {
			fmt.Fprint(os.Stderr,"fetch reading %s:%v \n",error)
			os.Exit(1)
		}
	}
	fmt.Printf("get all data:%s",sresult)
}