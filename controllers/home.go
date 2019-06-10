package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Index() {
	this.Data["Js"] = "<script src=\"static/js/echarts.min.js\"></script>"
	this.TplName = "home.tpl"
}