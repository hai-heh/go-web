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
	Teacher Teacher `gorm:"FOREIGNKEY:TeacherId;ASSOCIATION_FOREIGNKEY:ID"`
	Courses []Course `gorm:"many2many:student_course;FOREIGNKEY:Id;ASSOCIATION_FOREIGNKEY:Id"`
}
type Teacher struct {
	Id int
	Name string
}

type Course struct {
	Id int
	Name string
}

type StudentCourse struct {
	Id int
	StudentId int
	CourseId int
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(101.200.205.9:3306)/test3?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	db.LogMode(true)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var s Student
	db.Preload("Teacher").Preload("Courses").First(&s, 1)
	fmt.Println(s)
}
