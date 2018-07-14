package controllers

import "github.com/astaxie/beego"

type User struct {
	Id 			int			`form:"-"`
	Username 	string		`form:"username"`
	Age 		string		`form:"age"`
	Email 		string		`form:"email"`
}

type TestArgController struct {
	beego.Controller
}

/**
func (c *TestArgController) Get()  {
	// 127.0.0.1:8082/?id=10&name=zhangsan
	id := c.GetString("id")
	c.Ctx.WriteString(id)
	name := c.Input().Get("name")
	c.Ctx.WriteString(name)
}*/

func (c *TestArgController) Get()  {
	c.TplName = "index.tpl"
}

func (c *TestArgController) Post()  {
	u := User{}
	if err := c.ParseForm(&u); err != nil {

	}
	c.Ctx.WriteString("Username: " + u.Username + "  Age: " + u.Age + "  Email: " + u.Email)
}