package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	Id int
	Name string
	TeacherId int
	Teacher Teacher
}

type Teacher struct {
	Id int
	Name string
	Students []Student `gorm:"FOREIGNKEY:TeacherId;ASSOCIATION_FOREIGNKEY:Id"`
}

func main()  {
	db, err := gorm.Open("mysql", "root:123456@tcp(101.200.205.9:3306)/test3?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	db.LogMode(true)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	defer db.Close()
	var t Teacher
	db.Preload("Students").First(&t,1)
	fmt.Println(t)
}