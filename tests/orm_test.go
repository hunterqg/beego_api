package test

import "testing"
import (
	"beego_api/models"
	_ "beego_api/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path/filepath"
	"runtime"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	fmt.Println("------->", apppath)
	beego.TestBeegoInit(apppath)
}
func TestPerson(t *testing.T) {
	o := orm.NewOrm()
	o.Using("default")
	p := models.Person{
		Name: "qingao",
		Age:  22, Gender: 1,
	}
	p.CreateAt = time.Now()
	p.UpdateAt = time.Now()
	id, err := o.Insert(&p)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("ID:", id, "inserted")
	}
}
