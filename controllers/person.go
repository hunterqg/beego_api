package controllers

import (
	"beego_api/dao"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type PersonController struct {
	beego.Controller
}

func (this *PersonController) Index() {
	personDao := dao.PersonDao{}
	ps := personDao.GetAll()
	d := make(map[string]interface{})
	d["code"] = 1
	d["persons"] = ps
	this.Data["json"] = d
	this.ServeJSON()
}

func (this *PersonController) Get() {
	param := this.Ctx.Input.Param(":id")
	logs.Debug("---->", param)
	personDao := dao.PersonDao{}
	d := make(map[string]interface{})

	id, err := strconv.Atoi(param)
	if err != nil {
		d["code"] = -1
		d["msg"] = err.Error()
	} else {
		p ,e:= personDao.GetById(int64(id))

		if e!=nil {
			d["code"] = -1
			d["msg"] = e.Error()
		}else{
			d["code"] = 1
			d["person"] = p
		}
	}

	this.Data["json"] = d
	this.ServeJSON()

}
