package main

import "fmt"

type test interface {
	say(a int,b int) int
}

type sayfunc func(a int,b int) int

func (s sayfunc) say(a int,b int) int {
    return s(a,b)
}

func main() {
    var a int =5
    f:= func() {
    	a++
	}
	update(f)
    fmt.Println(a)

}

func update( f func())  {
	update2(f)
}

func update2(f func())  {
	f()
}

