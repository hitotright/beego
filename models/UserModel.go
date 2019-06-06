package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/hitotright/beego-admin/lib"
	"time"
)

type User struct {
	UserId        int     `orm:"auto;" description:"用户编号"`
	Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(2)" description:"用户名"`
	LoginName     string    `orm:"unique;size(32)" form:"LoginName" valid:"Required;MaxSize(20);MinSize(6)" description:"登录名"`
	Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)" description:"密码"`
	RePassword    string    `orm:"-" form:"RePassword" valid:"Required" description:"找回密码"`
	Email         string    `orm:"size(32)" form:"Email" valid:"Email" description:"邮箱"`
	Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)" description:"备注"`
	Status        int       `orm:"default(1)" form:"Status" valid:"Range(1,2)" description:"状态：1正常2禁止"`
	LastLoginTime time.Time `orm:"null;type(datetime)" form:"-" description:"最后登录时间"`
	CreateTime    time.Time `orm:"type(datetime);auto_now_add" description:"创建时间"`
	DpId          int     `orm:"default(1)" description:"部门编号"`
	PsId          int     `orm:"default(1)" description:"职位编号"`
}

func CheckLogin(login_name, password string) (user User, err error) {
	user = GetUserByLoginName(login_name)
	if (user.UserId == 0) {
		return user, errors.New("用户不存在")
	}
	if user.Password != lib.PwdHash(password) {
		return user, errors.New("密码错误")
	}
	return user, nil
}

func GetUserByLoginName(login_name string) (user User) {
	user = User{LoginName: login_name}
	o := orm.NewOrm()
	_ = o.Read(&user, "LoginName")
	return user
}

func init() {
	orm.RegisterModel(new(User))
}
