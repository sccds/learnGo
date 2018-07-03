// mysql add/insert data

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

var Db *sqlx.DB

func init()  {
	database, err := sqlx.Open("mysql", "sccds:123456@tcp(127.0.0.1:3306)/51ctogo")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	Db = database
}

func main()  {
	r, err := Db.Exec("INSERT INTO person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	fmt.Println("insert succ:", id)
}