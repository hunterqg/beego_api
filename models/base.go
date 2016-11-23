package models

import "time"

type base struct {
	Id int64
	CreateAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)"`
}
