package main


import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var age int

type Person struct {
	mux sync.Mutex
}

func (p *Person) AddAge() {
	defer wg.Done()
	p.mux.Lock()
	age++
	defer p.mux.Unlock()
}

func main() {
	p1 := Person{
		mux: sync.Mutex{},
	}
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go func() {
			p:=&p1;
			p.AddAge()
		}()
	}
	wg.Wait()
	fmt.Println(age)
}