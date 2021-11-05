
package main

import (
"fmt"
"sync"
)


var s sync.WaitGroup;
type Foo struct {
	FirstToSecond chan int
	SecondToThird chan int
	ThirdToEnd chan int
}

func NewFoo() *Foo {
	f:=Foo{}
	f.FirstToSecond=make(chan int,1)
	f.SecondToThird=make(chan int,1)
	f.ThirdToEnd=make(chan int,1)
	f.FirstToSecond<-1;
	return &f
}





func (this *Foo) First()  {
	defer s.Done()
	<-this.FirstToSecond
	fmt.Println("first")
	this.SecondToThird<-1
}

func (this *Foo) second()  {
	defer s.Done()
	<-this.SecondToThird
	fmt.Println("second")
	this.ThirdToEnd<-1
}

func (this *Foo) third()  {
	defer s.Done()
	<-this.ThirdToEnd
	fmt.Println("third")
}

func main()  {
	s.Add(3)
	f:=NewFoo()
	go f.second()
	go f.third()
	go f.First()
	s.Wait()
}
