package models

import "time"

type base struct {
	Id int64
	CreateAt time.Time
	UpdateAt time.Time
}
