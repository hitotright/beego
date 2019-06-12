package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	MenuId          int       `orm:"auto;" description:"编号"`
	PId        	    int       `orm:"int(11);default(0)" form:"PId" description:"上级菜单"`
	Name            string    `orm:"size(20);default('')" form:"Name" valid:"MaxSize(20)" description:"菜单名"`
	Url             string    `orm:"size(255);default('')" form:"Url" valid:"MaxSize(255)" description:"地址"`
	Icon            string    `orm:"size(20);default('')" form:"Icon" valid:"MaxSize(20)" description:"菜单图标"`
	IconColor       string    `orm:"size(20);default('')" form:"IconColor" valid:"MaxSize(20)" description:"菜单颜色"`
	Sort            int       `orm:"int(11);default(0)" form:"Sort" description:"排序"`
	Disabled        int8      `orm:"tinyint(1);default(1)" form:"Disabled" description:"1=可用，0=禁用"`
	PermissionId    int       `orm:"int(11);default(0)" form:"PermissionId" description:"权限ID"`
	TabId           string    `orm:"null;unique;size(255)" form:"TabId" valid:"MaxSize(255)" description:"菜单唯一标识"`
}

func (m *Menu) TableEngine() string {
    return "INNODB"
}

type menuTree struct {
	MenuId          int64
	PId        	    int64
	Name            string
	Url             string
	Icon            string
	IconColor       string
	TabId           string
	Children		[]menuTree
}

func GetMenu() ([]menuTree) {
	var cMenu, pMenu []orm.Params
	//第一层
	o := orm.NewOrm()
	menu := new(Menu)
	_, err := o.QueryTable(menu).Filter("PId", 0).Values(&pMenu)
	if err != nil {
		logs.Error(err)
		return []menuTree{}
	}
	var pMenuIndex = make(map[int64]int,len(pMenu))
	tree := make([]menuTree, len(pMenu))
	for i, p := range pMenu {
		pMenuIndex[p["MenuId"].(int64)] = i
		tree[i].MenuId = p["MenuId"].(int64)
		tree[i].PId = 0
		tree[i].Name = p["Name"].(string)
		tree[i].Url = p["Url"].(string)
		tree[i].Icon = p["Icon"].(string)
		tree[i].IconColor = p["IconColor"].(string)
		tree[i].TabId = p["TabId"].(string)
	}
	//第二层
	o = orm.NewOrm()
	_, err = o.QueryTable(menu).Filter("PId__gt", 0).OrderBy("PId").OrderBy("-Sort").Values(&cMenu)
	if err != nil {
		logs.Error(err)
		return tree
	}
	for _, m := range cMenu {
			index,ok := pMenuIndex[m["PId"].(int64)]
			if(ok == true){
				c := menuTree{m["MenuId"].(int64),m["PId"].(int64),m["Name"].(string),m["Url"].(string),
					m["Icon"].(string),m["IconColor"].(string),m["TabId"].(string),[]menuTree{}}
				tree[index].Children = append(tree[index].Children,c)
			}
	}
	return tree
}

func init() {
	orm.RegisterModel(new(Menu))
}