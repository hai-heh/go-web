package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	hash()
}

func  getCoon()  redis.Conn{
	c, err := redis.Dial("tcp", "101.200.205.9:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return nil
	}
	fmt.Println("redis conn success")
	return c
}

func string()  {
	coon := getCoon()
	defer coon.Close()
	_, err := coon.Do("set", "abc", 150)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	r, err:= redis.Int(coon.Do("get", "abc"))
	if(err!=nil){
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}

func hash()  {
	coon:=getCoon()
	defer coon.Close()

	_, err := coon.Do("hset", "abcd","hhh", 350)
	coon.Do("expire","abcd",10)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	r, err:= redis.Int(coon.Do("hget", "abcd","hhh"))
	if(err!=nil){
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
