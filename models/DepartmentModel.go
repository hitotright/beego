package models

import "github.com/astaxie/beego/orm"

type Department struct {
	DpId        int     `orm:"auto;" description:"部门编号"`
	PId        	int     `orm:"int(11)" form:"-" description:"部门编号"`
	DpName      int    `orm:"unique;size(32)" form:"DpName"  valid:"Required;MaxSize(20);MinSize(2)" description:"部门名"`
}

func (dp *Department) TableEngine() string {
    return "INNODB"
}

func init() {
	orm.RegisterModel(new(Department))
}