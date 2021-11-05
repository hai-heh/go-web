package main

import (
	"fmt"
)

func main()  {
	var nums [5]int=[5]int{1,2,3,4,5} // 建立数组
	var mySlice []int=nums[:] // 转化为切片

	mySlice1:=make([]int,5) // 通过make去初始化空间,相当于执行构造函数
	fmt.Println(mySlice1)

	add(&mySlice)
	fmt.Println(mySlice)
	deleteFirst(&mySlice)
	fmt.Println(mySlice)
	deleteLast(&mySlice)
	fmt.Println(mySlice)
}

// add 添加元素
func add(mySlice *[]int)  {
	*mySlice=append(*mySlice,2)
}

// deleteFirst 删除第一个元素
func deleteFirst(mySlice *[]int)  {
	*mySlice=append((*mySlice)[:0],(*mySlice)[1:]...)
}

// deleteLast 删除最后一个元素
func deleteLast(mySlice *[]int)  {
	*mySlice=append((*mySlice)[0:len(*mySlice)-1],(*mySlice)[len(*mySlice):]...)
}


