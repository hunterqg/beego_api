package models

import "time"

type Base struct {
	Id int64 `orm:"pk"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)"`
}
