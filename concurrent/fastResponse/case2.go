package main

import (
	"context"
	"fmt"
)

func main()  {
	stream:=make(chan int,1)
	ctx, cancel := context.WithCancel(context.Background())
	var find string="a"
	var nums []string=[]string{"a","aa","aaa","aaaa","aaaaa","aaaaaa"}
	var length int= len(nums)/10
	for i:=0;i<length;i++{
		go func() {
			for j:=i*10;j<(i+1)*10;j++{
				select {
				case <-ctx.Done():
					return
				}
				if(nums[j]==find){
					stream<-j
				}
			}
		}()
	}
	index:=<-stream
	fmt.Println(index)
	cancel()

}
