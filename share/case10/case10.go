package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	course Course // 如果是指针，这个结构体大小为多少呢
	c int64
}
type Course struct {
	a int64 // 8字节
	b int64 // 8字节
}
func main()  {
	fmt.Println(unsafe.Sizeof(Student{}))
	fmt.Println(unsafe.Sizeof(Course{}))
}
