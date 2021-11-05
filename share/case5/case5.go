package main

import "fmt"

func main()  {
	hashMap:=make(map[int]string,10)
	hashMap[1]="小明"
	hashMap[2]="小红"
	hashMap[3]="小刚" // put一个节点到哈希表中
	fmt.Println(hashMap[1]) // get哈希表某一个值
	for index,v:=range hashMap { // 遍历哈希表的值
		if index != 1 {
		fmt.Println(v)
		}
	}

}
