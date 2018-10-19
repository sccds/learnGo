// mysql update

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
	_, err := Db.Exec("UPDATE person set username = ? WHERE user_id = ?", "stu0003", 1)
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}

}