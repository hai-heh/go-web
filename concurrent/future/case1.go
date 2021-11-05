package main

import (
	"fmt"
	"time"
)

type future struct {
	 stream chan int
}

// 做业务逻辑
func (this *future) set(value int) {
	time.Sleep(time.Second*2)
	this.stream<-value
}

func (this *future) get() int {
	return <-this.stream
}

func main()  {
	f:=future{stream:make(chan int,1)}
	go f.set(3)
	fmt.Println(f.get())
}
