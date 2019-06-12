package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/lib"
	"github.com/hitotright/beego-admin/models"
	_ "github.com/hitotright/beego-admin/routers"
	"io/ioutil"
	"os"
	"strings"
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
	models.Init()
	name := "default"
	force := true
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("insert add user ...")
	u := new(models.User)
	u.Username = "admin"
	u.LoginName = "admin"
	u.Password = lib.Pwdhash("123456")
	o := orm.NewOrm()
	o.Insert(u)
	fmt.Println("insert user end")
	fmt.Println("exec init.sql ...")
	f, err := os.Open("init.sql")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	sql,_ := ioutil.ReadAll(f)
	sql_slice :=strings.Split(string(sql),";")
	for _,s := range sql_slice  {
		s = strings.ReplaceAll(s,"\n","")
		s = strings.ReplaceAll(s,"\r","")
		if(s != ""){
			fmt.Println(s)
			_, err :=o.Raw(s).Exec()
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}
	}
	fmt.Printf("exec init.sql complete：%d\n",len(sql_slice))
}
