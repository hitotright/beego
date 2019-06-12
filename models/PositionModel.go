package models

import "github.com/astaxie/beego/orm"

type Position struct {
	PsId        int     `orm:"auto;" description:"职位编号"`
	PId        	int     `orm:"int(11)" form:"-" description:"职位编号"`
	PsName      string    `orm:"unique;size(32)" form:"PsName"  valid:"Required;MaxSize(20);MinSize(2)" description:"职位名"`
}

func (ps *Position) TableEngine() string {
    return "INNODB"
}

func init() {
	orm.RegisterModel(new(Position))
}