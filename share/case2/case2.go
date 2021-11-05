package main

import "fmt"


type Course struct {
	courseName string
}
type Student struct {
	name string
	age int
	course Course

}

// 赋值指针不一样的
func main()  {
	s:=Student{name: "小明",age: 32,course: Course{courseName: "数学",
	}}
	s1:=s
	fmt.Printf("s指针为%p\n",&s)
	fmt.Printf("s1指针为%p\n",&s1)
	fmt.Printf("s的course指针为%p\n",&(s.course))
	fmt.Printf("s1的course指针为%p\n",&(s1.course))
}

func copy(s2 Student)  {
	fmt.Printf("s2指针为%p\n",&s2) // 0xc0000044c0
}