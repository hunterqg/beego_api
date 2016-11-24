package utils

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
)
var Cache cache.Cache

func init() {
	logs.Info("Init db connection ...")
	Cache,_ = cache.NewCache("memory",`interval:60`)
}
