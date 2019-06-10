package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Index() {
	this.TplName = "home.tpl"
}