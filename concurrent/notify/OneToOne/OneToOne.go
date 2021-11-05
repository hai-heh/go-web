package main

import (
	"fmt"
	"sort"
)

func main()  {
	var num []int=[]int{5,4,3,2,1}
	stream:=make(chan struct{})
	go func() {
		sort.Slice(num, func(i, j int) bool {
			return num[i]<num[j]
		})
		stream<- struct {}{}
	}()
	<-stream
	fmt.Println(num)
}
