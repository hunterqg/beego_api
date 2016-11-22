package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	_ "beego_api/utils"
)

type Person struct {
	base
	Name   string `orm:"size(50)"`
	Age    int    `orm:"int(4)"`
	Gender int    `orm:"int(1)"`
}


func init() {
	logs.Info("Init the models : person")
	orm.RegisterModel(new(Person))

	orm.RunSyncdb("default", false, true)

}