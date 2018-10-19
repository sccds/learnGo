// mysql select

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

var Db *sqlx.DB

// 创建结构体，需要和DB中的schema绑定
type Person struct {
	UserId int `db:"user_id"`
	UserName string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}


func init()  {
	database, err := sqlx.Open("mysql", "sccds:123456@tcp(127.0.0.1:3306)/51ctogo")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	Db = database
}

func main()  {
	var person []Person
	err := Db.Select(&person, "SELECT user_id, username, sex, email FROM person")
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	fmt.Println("select succ:", person)
}
