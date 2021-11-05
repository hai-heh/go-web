package main

import "fmt"

func f() func()   {
	var sum int=2
	return func() { // 这个匿名函数有一个外部全局变量sum
		sum+=2
		fmt.Println(sum)
	}
}


func main() {
	f1:=f()
	f1() // 4
	f1() // 6
	f2:=f()
	f2() // 4
	f2() // 6
}
