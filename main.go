package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/lib"
	"github.com/hitotright/beego-admin/models"
	_"github.com/hitotright/beego-admin/routers"
	"os"
)

func main() {
	args := os.Args
	for _, v := range args {
		if v == "-init" {
			dbInit()
			os.Exit(0)
		}
	}
	beego.Run()
}

func dbInit()  {
	models.Connect()
	name := "default"
	force := true
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("insert add user ...")
	u := new(models.User)
	u.Username = "admin"
	u.LoginName = "admin"
	u.Password = lib.Pwdhash("123456")
	o := orm.NewOrm()
	o.Insert(u)
	fmt.Println("insert user end")
}
