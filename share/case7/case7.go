package main

import "fmt"


// 恐慌回复机制
func main()  {
	defer func() { // 该方法进栈
		fmt.Println("正常退出")
	}()
	fmt.Println("嗨！")
	defer func() {  // 该方法进栈
		v := recover() // 恢复恐慌机制
		fmt.Println("恐慌被恢复了：", v)
	}()
	panic("拜拜！") // 产生一个恐慌
	fmt.Println("执行不到这里")
}
