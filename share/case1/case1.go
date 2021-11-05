package main

import "fmt"




type student struct {
	name string
	age int
}

// 赋值指针不一样的
func main()  {


	s:=student{name: "小明",age: 32}
	s1:=s
	fmt.Printf("s指针为%p\n",&s) // 0xc000004460
	fmt.Printf("s1指针为%p\n",&s1) // 0xc000004480
	copy(s)
}

func copy(s2 student)  {
	fmt.Printf("s2指针为%p\n",&s2) // 0xc0000044c0
}
