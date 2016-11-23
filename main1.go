package main

import (
	//"beego_api/models"
	_ "beego_api/utils"
	"github.com/astaxie/beego/orm"
	"fmt"
	//"time"
	"beego_api/models"
	"time"
)

func main1() {
	o := orm.NewOrm()
	fmt.Println(o)
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
	/*cnt,e := o.Delete(&p)
	if e != nil {
		fmt.Println("Error:", e.Error())
	}else{
		fmt.Println("deleted ", cnt, "persons")
	}

	p1 := models.Person{}
	p1.Id = 15
	e = o.Read(&p1)
	if e!= nil {
		fmt.Println("Error:", e.Error())
	}else{
		fmt.Println("Record:",p1)
	}
	*/

}
