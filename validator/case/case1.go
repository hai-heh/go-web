package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Age  int    `validate:"min=1,max=100"`
    Name string
}

func main()  {
	validate:=validator.New()
	u:=User{Age: 126,Name: "小明"}
	err := validate.Struct(u)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(u)
}
