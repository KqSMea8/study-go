package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

// person
type Person struct {
	ID     uint32
	Sex    bool
	Name   string
	Weight float32
	Friend *Friends
	Books  []string
}

// friends
type Friends struct {
	Age  int
	Name string
}

func ReadConf(fname string) (p *Person, err error) {
	var (
		fp       *os.File
		fcontent []byte
	)
	p = new(Person) // &Person{}
	if fp, err = os.Open(fname); err != nil {
		fmt.Println("open error ", err)
		return
	}

	if fcontent, err = ioutil.ReadAll(fp); err != nil {
		fmt.Println("ReadAll error ", err)
		return
	}

	if err = toml.Unmarshal(fcontent, p); err != nil {
		fmt.Println("toml.Unmarshal error ", err)
		return
	}
	return
}