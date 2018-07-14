package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController) Get()  {
	c.Ctx.WriteString("这是第一个beego控制器Get方法")
}

func (c *TestController) Post()  {
	c.Ctx.WriteString("这是第一个beego控制器POST方法")
}

