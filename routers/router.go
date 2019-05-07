package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin"
	"quickstart/controllers"
)

func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/user",&controllers.UserController{})
}
