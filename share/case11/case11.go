package main

import (
	"fmt"
	"unsafe"
)

type Student1 struct {
	a byte
	b int32 // a+b=5 只需要填充3个字节
	c int64 // 8个字节不用填充
}

type Student2 struct {
	a byte // 填充7字节
	c int64
	b int32 // 填充4字节
}

func main()  {
	fmt.Println(unsafe.Sizeof(Student1{}))
	fmt.Println(unsafe.Sizeof(Student2{}))
}