package dao

import (
	"beego_api/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)
type PersonDao struct {}


func (this *PersonDao) GetById(id int64) ( models.Person, error ) {
	logs.Debug("start get person by id...")
	o := orm.NewOrm()
	p := models.Person{}
	p.Id = id

	error := o.Read(&p)
	if error != nil {
		logs.Info("---->Person not find : ",error.Error())
		return p,error
	}
	return p,nil

}

func (this *PersonDao) GetAll()([] *models.Person) {
	var persons []*models.Person
	o := orm.NewOrm()
	_,err := o.QueryTable("person").All(&persons)
	if err !=nil {
		logs.Info("Query Persion Failed : ",err.Error())
	}
	return persons

}
