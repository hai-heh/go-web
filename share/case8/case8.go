package main

import "fmt"

func main()  {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println(err)
		}
	}()
	defer func() {
		a:=1/0
		fmt.Println(a)
	}()
	defer panic(1)
	panic(0)
}
