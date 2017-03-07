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

var o = orm.NewOrm()

func ( *PersonDao) GetById(id int64) ( models.Person, error ) {
	logs.Debug("start get person by id...")
	key := generateKey(id)
	p := models.Person{}
	if(utils.Cache.IsExist(key)) {
		logs.Debug("get person from cache")
		p = utils.Cache.Get(key).(models.Person)
		return p , nil
	}
	p.Id = id

	error := o.Read(&p)
	if error != nil {
		logs.Info("---->Person not find : ",error.Error())
		return p,error
	}
	utils.Cache.Put(key,p,3*time.Minute)
	return p,nil

}

func ( *PersonDao) GetAll()([] *models.Person) {
	var persons []*models.Person
	_,err := o.QueryTable("person")/*.Limit(50,100)*/.All(&persons)
	if err !=nil {
		logs.Info("Query Persion Failed : ",err.Error())
	}
	return persons

}

func ( *PersonDao)AddOne(p *models.Person) (int64,error) {
	return o.Insert(p)
}

func (*PersonDao) DeleteById(id int64) (int64,error){
	p := models.Person{}
	p.Id = id
	return o.Delete(&p)
}

func generateKey(id int64) string {
	return fmt.Sprintf("Person_GetById_%d",id)
}