package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Data["name"] = "tom"
	this.Data["age"] = 23
	this.TplName = "user.tpl"
}

func (this * UserController) Index()  {

}

func (this * UserController) Add()  {

}

func (this * UserController) Edit()  {

}

func (this * UserController) Delete()  {

}


