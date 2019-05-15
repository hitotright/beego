package controllers

import "github.com/hitotright/beego-admin/models"

type LoginController struct {
	BaseController
}

//登录
func (this *LoginController) Login(){
	if this.IsAjax() {
		login_name := this.GetString("login_name")
		password := this.GetString("password")
		user, err := models.CheckLogin(login_name, password)
		if err == nil {
			this.SetSession("user_info", user)
			this.Rsp(true, "登录成功")
			return
		} else {
			this.Rsp(false, err.Error())
			return
		}

	}
	user_info := this.GetSession("user_info")
	if user_info != nil {
		this.Ctx.Redirect(302, "/")
	}
	this.TplName = "login.tpl"
}