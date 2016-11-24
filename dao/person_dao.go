package dao

import (
	"beego_api/models"
	"beego_api/utils"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"fmt"
	"time"
)
type PersonDao struct {}


func (this *PersonDao) GetById(id int64) ( models.Person, error ) {
	logs.Debug("start get person by id...")
	key := generateKey(id)
	p := models.Person{}
	if(utils.Cache.IsExist(key)) {
		logs.Debug("get person from cache")
		p = utils.Cache.Get(key).(models.Person)
		return p , nil
	}
	o := orm.NewOrm()
	p.Id = id

	error := o.Read(&p)
	if error != nil {
		logs.Info("---->Person not find : ",error.Error())
		return p,error
	}
	utils.Cache.Put(key,p,3*time.Minute)
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

func generateKey(id int64) string {
	return fmt.Sprintf("Person_GetById_%d",id)
}