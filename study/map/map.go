package main

import (
	"fmt"
	"strconv"
)

func main()  {
	m:=make(map[string]int,7)
	for i:=0;i<1;i++{
		m[strconv.Itoa(i)]=23
	}
	a:=m["1"]
	fmt.Println(a)
}



