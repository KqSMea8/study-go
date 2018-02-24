package main

import (
	//"os"
	"fmt"
	//"net/http"
)

func main() {
	/*var s,sep stgolintgolintgolintring
	for i:=1;i<len(os.Args);i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	*/
	a := [...]int{9, 1, 2, 3, 4, 5,10}
	b := [...]string{"a","ba","cc"}
	mmap := make(map[string]int) //map[string]int{}
	mmap["xiaojh"]=1
	mmap["sex"]=2
	mmap["age"]=22


	//mrang(mmap)
	fmt.Println("原始：",a[:])
	reverse(a[:])
	fmt.Println("反转：",a[:])
	msort(a[:])
	fmt.Println("冒泡.倒：",a[:])
	mrang(b[:])
	selectFind(a[:])
	fmt.Println("选择",a[:])

	c := [...] int {1,2,3,4,5,6,7,8,9,10}
	findx := erFind(c[:],10)
	fmt.Println("二分查：",c[:]," 10的位置：",findx)
}

//反转
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//大冒泡
func msort(s []int) {
	tmp := 0
	for k := 0; k < len(s)-1; k++ {
		for i := 0; i < len(s)-k-1; i++ {
			if s[i] < s[i+1] {
				tmp = s[i]
				s[i] = s[i+1]
				s[i+1] = tmp
			}
		}
	}
}

//二分法
func erFind(s [] int,xfind int) int {
	start,xclen := 0,len(s)
	for start < xclen {
		mid := (start+xclen) >> 1
		if s[mid]<xfind {
			start = start+1
		} else if s[mid] > xfind {
			xclen = mid-1
		} else {
			//return s[mid]
			return mid
		}
	}
	return -1
}

//选择排序
func selectFind(s [] int)  {
	xlen := len(s)
	var tmp, k int
	for i := 0; i < xlen; i++ {
		k = i
		for j := i + 1; j < xlen; j++ {
			if s[j] < s[k] {
				k = j
			}
		}
		if i != k {
			tmp = s[i]
			s[i] = s[k]
			s[k] = tmp
		}
	}
	//return s
}

func mrang(s []string) {
	for line,n:=range s {
		if line==0 {
			//continue
		}
		fmt.Printf("%s=>%s\n",line,n)
	}
}
