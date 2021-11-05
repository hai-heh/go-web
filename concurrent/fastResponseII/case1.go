package main

import (
	"context"
	"fmt"
	"sync"
)


func Min(a int,b int) int{
   if(a<b){
   	return a
   }else{
   	return b
   }
}
func main()  {
	stream:=make(chan int,1)
	ctx, cancel := context.WithCancel(context.Background())
	var find string="1"
	var s string=""
	var nums []string=[]string{}
	for k:=0;k<100;k++{
		nums=append(nums,s)
		s=s+"a"
	}
	var length int= (len(nums)/10)+1
	var group sync.WaitGroup
	group.Add(length)
	count:=0
	lock:=sync.Mutex{}
	taskIndex:=0
	stream1:=make(chan struct{},1)
	for taskIndex<length{
		go func() {
			stream1<-struct{}{}
			i:=taskIndex
			taskIndex++;
			for j:=i*10;j<Min((i+1)*10,len(nums));j++{
				if(nums[j]==find){
					stream<-j
					return
				}
				select {
				case <-ctx.Done():
					fmt.Println("挂了")
					return
				default:{}
				}
			}
			lock.Lock()
			defer lock.Unlock()
			count++
		}()
		<-stream1
	}
	var index int=-1
	for {
		select {
		case index = <-stream:{
			fmt.Println(index)
			cancel()
			return
		}
		default:
			if(count==length){
				fmt.Println("找不到")
				return
			}
		}
	}
}