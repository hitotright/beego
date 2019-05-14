package models


type Department struct {
	DpId        int64     `orm:"auto;rel(fk)" description:"部门编号"`
	PId        	int64     `orm:"int(11)" form:"-" description:"部门编号"`
	DpName      string    `orm:"unique;size(32)" form:"DpName"  valid:"Required;MaxSize(20);MinSize(2)" description:"部门名"`
}
