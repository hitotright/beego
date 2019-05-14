package models


type Position struct {
	PsId        int64     `orm:"auto;rel(fk)" description:"职位编号"`
	PId        	int64     `orm:"int(11)" form:"-" description:"职位编号"`
	PsName      string    `orm:"unique;size(32)" form:"PsName"  valid:"Required;MaxSize(20);MinSize(2)" description:"职位名"`
}