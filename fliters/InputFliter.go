package fliters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"encoding/json"
)

var InputFliter1 = func (ctx *context.Context) {
	logs.Debug("------1--->",ctx.Input.Params())
	for k,v := range ctx.Input.Params() {
		logs.Debug("-----1---->%v => %v",k,v)
	}
}
var jsonInputFliter = func (ctx *context.Context) {
	for k,v := range ctx.Input.Params() {
		logs.Info("Params:{ %v => %v }",k,v)
	}
	if ctx.Request.Method == "POST"{
		var data = make(map[string]interface{})
		err := json.Unmarshal(ctx.Input.RequestBody,&data)
		if err != nil {
			ctx.Abort(500,"wrong input json")
		}
		logs.Info("Post Body : ",data)
		ctx.Input.SetData("parsedMap",data)
	}
}

func init() {
	logs.Info("Register fliters")
	//beego.InsertFilter("/*",beego.BeforeRouter,InputFliter1)
	beego.InsertFilter("/*",beego.BeforeExec,jsonInputFliter)
}