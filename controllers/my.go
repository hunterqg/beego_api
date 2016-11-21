package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyController struct {
	beego.Controller
}

func (this *MyController) Get() {
	d := make(map[string] interface{})
	d["code"] = 1
	d["msg"]= "hello"
	this.Data["json"]= d

	logs.Debug("input:",this.GetString("k"))
	this.Ctx.WriteString("------------>")
	this.ServeJSON()
}

func (this *MyController)Lookup() {
	d := make(map[string] interface{})
	d["msg"]= "lookup"
	this.Data["json"]= d
	//logs.Debug(this.Ctx.Input.Params()["0"])
	//logs.Debug(this.Ctx.Input.Params()["1"])
	//logs.Debug(this.Ctx.Input.Params()[":splat"])
	for k,v := range this.Ctx.Input.Params() {
		logs.Debug("%v:%v",k,v)
	}

	this.ServeJSON()
}
