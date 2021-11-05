package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
`gorm:"foreignkey:Name;association_foreignkey:Name"`
*/
type Student struct {
	Id int
	Name string
	TeacherId int
	Teacher Teacher `gorm:"FOREIGNKEY:TeacherId;association_foreignkey:Id"`

}

type Teacher struct {
	Id int
	Name string
}

func main()  {
	fmt.Println()
	db, err := gorm.Open("mysql", "root:123456@tcp(101.200.205.9:3306)/test3?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	db.LogMode(true)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	defer db.Close()
	var s Student
	db.Preload("Teacher").First(&s,1)
	fmt.Println(s)
	
}

