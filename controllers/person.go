package controllers

import (
	"beego_api/dao"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"beego_api/models"
	//"time"
)

type PersonController struct {
	beego.Controller
}

var pDao = dao.PersonDao{}

func (this *PersonController) Index() {
	ps := pDao.GetAll()
	d := make(map[string]interface{})
	d["code"] = 1
	d["persons"] = ps
	this.Data["json"] = d
	this.ServeJSON()
}

func (this *PersonController) Get() {
	param := this.Ctx.Input.Param(":id")
	logs.Debug("---->", param)
	d := make(map[string]interface{})

	id, err := strconv.Atoi(param)
	if err != nil {
		d["code"] = -1
		d["msg"] = err.Error()
	} else {
		p, e := pDao.GetById(int64(id))

		if e != nil {
			d["code"] = -1
			d["msg"] = e.Error()
		} else {
			d["code"] = 1
			d["person"] = p
		}
	}

	this.Data["json"] = d
	this.ServeJSON()

}

func (this *PersonController) Post() {
	d := this.Ctx.Input.GetData("parsedMap")
	ret := make(map[string]interface{})
	if d == nil {

		ret["msg"] = "no input"
	} else {
		d1 := d.(map[string]interface{})
		p := models.Person{}
		p.Age = uint8(d1["Age"].(float64))
		p.Gender = uint8(d1["Gender"].(float64))
		p.Name = d1["Name"].(string)
		//p.CreateAt = time.Now()
		//p.UpdateAt = time.Now()
		id, err := pDao.AddOne(&p)
		if err != nil {
			ret["msg"] = "add to db failed:" + err.Error()
		} else {
			ret["id"] = id
			ret["msg"] = "ok"
		}
	}
	this.Data["json"] = ret
	this.ServeJSON()
}

func (this *PersonController) Delete() {
	param := this.Ctx.Input.Param(":id")
	d := make(map[string]interface{})
	id, err := strconv.Atoi(param)
	if err != nil {
		d["code"] = -1
		d["msg"] = err.Error()
	} else {
		_, e := pDao.DeleteById(int64(id))

		if e != nil {
			d["code"] = -1
			d["msg"] = e.Error()
		} else {
			d["code"] = 1
			d["msg"] = "ok"
		}
	}

	this.Data["json"] = d
	this.ServeJSON()
}