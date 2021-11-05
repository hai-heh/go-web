package main

import (
	"fmt"
	"time"
)

type timer struct {

}

func after(stream chan<- struct{})  {
	time.Sleep(time.Second)
	stream<- struct{}{}
}


func main()  {
	stream:=make(chan struct{})
	fmt.Println("hi")
	go after(stream)
	<-stream
	fmt.Println("hello")
	go after(stream)
	<-stream
	fmt.Println("你好")

}
