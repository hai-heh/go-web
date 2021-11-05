package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


type DbWorker struct {
	//mysql data source name
	Dsn string
}

type Student struct {
	Id int `db:"id"`
	Name string `db:"name"`
}

func main() {
	Select()
}

func Select()  {
	dsn:="root:123456@tcp(101.200.205.9:3306)/test3"
	coon, err := sqlx.Open("mysql", dsn)
	if(err!=nil){
		fmt.Println(err)
	}
	var s []Student
	err= coon.Select(&s, "select * from student where id>?", 0)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	for _,v:= range(s){
		fmt.Println(v)
	}
}


func Insert() {
	dbw := DbWorker{
		Dsn: "root:123456@tcp(101.200.205.9:3306)/test3",
	}
	db, err := sql.Open("mysql",
		dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO student(NAME) VALUE('aaa')")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := stmt.Exec()
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}