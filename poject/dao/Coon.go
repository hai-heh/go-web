package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func getCoon() *gorm.DB  {
	db, err := gorm.Open("mysql", "root:123456@tcp(101.200.205.9:3306)/test3?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if(err!=nil){
		fmt.Println(err)
	}
	return db
}

func close(db *gorm.DB)  {
	db.Close()
}
