package router

import (
	"SecKill/SecWeb/controller/activity"
	"SecKill/SecWeb/controller/product"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &product.ProductController{}, "*:ListProduct")
	beego.Router("/product/list", &product.ProductController{}, "*:ListProduct")
	beego.Router("/product/create", &product.ProductController{}, "*:CreateProduct")
	beego.Router("/product/submit", &product.ProductController{}, "*:SubmitProduct")
	beego.Router("/activity/list", &activity.ActivityController{}, "*:ListActivity")
	beego.Router("/activity/create", &activity.ActivityController{}, "*:CreateActivity")
	beego.Router("/activity/submit", &activity.ActivityController{}, "*:SubmitActivity")
}
