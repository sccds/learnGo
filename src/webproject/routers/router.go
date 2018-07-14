package routers

import (
	"webproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Get;post:Test")
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/test/arg", &controllers.TestArgController{}, "get:Get;post:Post")
	beego.Router("/test/orm", &controllers.TestModelController{}, "get:Get;post:Post")
	beego.Router("/test/httplib", &controllers.TestHttpLibController{}, "get:Get;post:Post")
}
