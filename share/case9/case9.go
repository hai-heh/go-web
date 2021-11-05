package main

import "fmt"

func f()  {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("重启一个协程处理")
			go f()
			select {}
		}
	}()
	panic("制造一个恐慌")
}

func main()  {
	f()
}
