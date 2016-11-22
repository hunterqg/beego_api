package utils
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/url"
)
func init() {
	logs.Info("Init db connection ...")
	orm.Debug = true
	initDB()
}


func initDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s",
		beego.AppConfig.String("db_user"),
		beego.AppConfig.String("db_password"),
		beego.AppConfig.String("db_host"),
		beego.AppConfig.String("db_port"),
		beego.AppConfig.String("db_name"),
		url.QueryEscape("Asia/Shanghai"),
	)

	logs.Debug(dbUrl)
	orm.RegisterDataBase(
		"default",
		"mysql",
		dbUrl,
		beego.AppConfig.DefaultInt("db_max_idle_conn", 20),
		beego.AppConfig.DefaultInt("db_max_open_conn", 20),
	)
}