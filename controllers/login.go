package controllers

type LoginController struct {
	BaseController
}

//登录
func (this *LoginController) Login(){
	if this.IsAjax() {
		login_name := this.GetString("login_name")
		password := this.GetString("password")
		user, err := CheckLogin(login_name, password)
		if err == nil {
			this.SetSession("user_info", user)
			access_list, _ := GetAccessList(user.Id)
			this.SetSession("access_list", access_list)
			this.Rsp(true, "登录成功")
			return
		} else {
			this.Rsp(false, err.Error())
			return
		}

	}
	user_info := this.GetSession("user_info")
	if user_info != nil {
		this.Ctx.Redirect(302, "/public/index")
	}
	this.TplName = "login.tpl"
}