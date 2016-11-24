package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	_ "beego_api/utils"
)

type Person struct {
	Base
	Name   string `orm:"size(50)"`
	Age    uint8
	Gender uint8
}


func init() {
	logs.Info("Init the models : person")
	orm.RegisterModel(new(Person))

	orm.RunSyncdb("default", false, true)

}
