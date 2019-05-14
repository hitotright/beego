package routers

import (
	"github.com/astaxie/beego"
	"github.com/hitotright/beego-admin/controllers"
)

func init() {
	beego.Router("/login",&controllers.LoginController{},"*:Login")
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/user",&controllers.UserController{})
}
