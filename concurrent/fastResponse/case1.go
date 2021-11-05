package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doSoming(stream chan int)  {
	a,b:=rand.Int31(),(rand.Intn(10)+1)
	time.Sleep(time.Second*time.Duration(b))
	stream<-int(a)
}

func main()  {
	startTime:=time.Now()
	stream:=make(chan int,5)
	for i:=0;i<5;i++{
        go doSoming(stream)
	}
	ans:=<-stream
	fmt.Println(time.Since(startTime))
	fmt.Println(ans)
}
