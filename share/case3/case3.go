
package main

import "fmt"


type Course struct {
	courseName string
}
type Student struct {
	course *Course
	name string
	age int
}

// 赋值指针不一样的
func main()  {
	s:=Student{name: "小明",age: 32,course: &Course{courseName: "数学",
	}}
	s1:=s
	fmt.Printf("s指针为%p\n",&s)
	fmt.Printf("s1指针为%p\n",&s1)


	fmt.Println()
	fmt.Printf("s的course指针为%p\n",s.course)
	fmt.Printf("s的course指针为%p\n",s.course)


	fmt.Println()
	fmt.Printf("s的course指针的指针为%p\n",&(s.course))
	fmt.Printf("s1的course指针的指针为%p\n",&(s1.course))


	fmt.Println()
	fmt.Printf("s指针为%p\n",&s)
	fmt.Printf("s的course指针的指针为%p\n",&(s.course)) // 思考一下如果结构体Student中的course位置改变答案一样吗
}

