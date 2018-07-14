package controllers

import (
	"github.com/astaxie/beego"
	"webproject/models"
)


type TestModelController struct {
	beego.Controller
}


func (c *TestModelController) Get() {
	//user := models.UserInfo{Username:"wangwu", Password:"123456"}
	//models.AddUser(&user)
	//c.Ctx.WriteString("call model success")

	var users []models.UserInfo
	models.ReadUserInfo(&users)
	c.Data["Title"] = "zhangsan"
	c.Data["Users"] = users
	c.Data["len"] = len(users)
	c.TplName = "test.tpl"
}