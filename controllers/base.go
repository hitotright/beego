package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/hitotright/beego-admin/models"
	"strconv"
	"time"
)

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
	this.Data["Menus"] = models.GetMenu()
}

func (this *BaseController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}

func init()  {
	date := time.Now().Format("20060102")
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/`+date+`/console.log",
"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	log_level,_ :=strconv.Atoi(beego.AppConfig.String("log_level"))
	logs.SetLevel(log_level)
	logs.SetLogFuncCall(true)
	logs.Async()
}