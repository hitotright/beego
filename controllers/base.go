package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare()  {
	this.Layout = "common/base.tpl"
	this.Data["Title"] = beego.AppConfig.String("title")
	this.Data["Keywords"] = beego.AppConfig.String("keywords")
	this.Data["Description"] = beego.AppConfig.String("description")
	this.Data["Css"] = ""
	this.Data["Js"] = ""
}

func (this *BaseController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}