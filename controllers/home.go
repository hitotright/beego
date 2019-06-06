package controllers

import (
	"github.com/hitotright/beego-admin/models"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Index() {
	this.Data["Menus"] = models.GetMenu()
	this.TplName = "home.tpl"
}