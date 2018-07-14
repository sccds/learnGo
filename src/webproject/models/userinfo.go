package models


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

var (
	db orm.Ormer
)

type UserInfo struct {
	Id int64
	Username string
	Password string
}


func init()  {
	orm.Debug = true  // 是否开启调试模式，调试模式下，会打印出SQL语句
	orm.RegisterDataBase("default", "mysql", "sccds:123456@/51ctogo?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))  // 创建一个user_info表 UserInfo
	db = orm.NewOrm()
}

// 添加一个user对象
func AddUser(user_info *UserInfo) (int64, error) {
	id, err := db.Insert(user_info)
	return id, err
}


// 读取一个对象
func ReadUserInfo(users *[]UserInfo) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user_info")
	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}
