package main

import (
	_ "beego_api/routers"
	_ "beego_api/fliters"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.Info("api is starting ...")
	logs.Info("Conf:",beego.BConfig)
	logs.Info("mysql:",beego.AppConfig.String("dev::mysql_user"),
		beego.AppConfig.String("dev::mysql_passwd"),
		beego.AppConfig.String("dev::mysql_dbName"),
	)
	ary := beego.AppConfig.Strings("ary_test");
	for _,x := range  ary{

		logs.Info(len(ary),x)
	}
	//beego.SetLevel(beego.LevelInformational)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		logs.Info("Running in DEV mode")
	}
	/*
	beego.SetStaticPath("/qingao","/home/qingao")
	beego.BConfig.WebConfig.DirectoryIndex = true
	*/
	beego.Run()
}
