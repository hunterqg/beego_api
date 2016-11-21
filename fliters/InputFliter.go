package fliters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)

var InputFliter1 = func (ctx *context.Context) {
	logs.Debug("------1--->",ctx.Input.Params())
	for k,v := range ctx.Input.Params() {
		logs.Debug("-----1---->%v => %v",k,v)
	}
}
var InputFliter2 = func (ctx *context.Context) {
	logs.Debug("-------2-->",ctx.Request.URL)
	for k,v := range ctx.Input.Params() {
		logs.Debug("----2----->%v => %v",k,v)
	}
}

func init() {
	logs.Info("Register fliters")
	beego.InsertFilter("/*",beego.BeforeRouter,InputFliter1)
	beego.InsertFilter("/*",beego.BeforeRouter,InputFliter2)
}