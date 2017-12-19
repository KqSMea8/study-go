package main

import (
	"bufio"
	"os"
	"strings"
	"io"
	"fmt"
	"time"
)

//读
func readfile(fileName string,handle func(string)) error {
	f,err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line,err := buf.ReadString('\n')
		//line = strings.TrimSpace(line)
		line = strings.TrimRight(line,"")
		handle(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

//写
func fwrite(fileName string,xcontent string) error {
	f,err := os.OpenFile(fileName,os.O_WRONLY|os.O_APPEND|os.O_CREATE,0666)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(xcontent)
	return nil
}

func printin (line string) {
	fmt.Printf(line)
}
func main() {
	readfile("mytext2.txt",printin)
	content := fmt.Sprintf("write as %s\n",time.Now())
	fwrite("mytext2.txt",content)
}
