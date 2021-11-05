package main

import (
	"fmt"
	"sync"
)

type Lock struct {

	stream chan struct{}
}

var group sync.WaitGroup
var sum int

func NewLock() *Lock {
	return &Lock{stream: make(chan struct{},1)}
}

func (this *Lock) Unlock()  {
	<-this.stream
}

func (this *Lock) Lock() {
	this.stream<- struct{}{}
}

func main()  {
	group.Add(10000)
	lock:=NewLock()
	for i:=0;i<10000;i++{
		go func() {
			lock.Lock()
			defer func() {
				group.Done()
				lock.Unlock()
			}()
			sum++
		}()
	}
	group.Wait()
	fmt.Println(sum)
}
