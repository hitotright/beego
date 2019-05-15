package routers

import (
	"github.com/astaxie/beego"
	"github.com/hitotright/beego-admin/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
}
