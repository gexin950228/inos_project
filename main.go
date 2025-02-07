package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "inos_project/models/user"
	_ "inos_project/routers"
)

func main() {
	beego.SetStaticPath("/static", "static")
	beego.SetViewsPath("views")
	err := logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/app.log", "separate": ["emergency", "critical", "error"]}`)
	if err != nil {
		logs.Emergency(err.Error())
	}
	orm.RunCommand()
	beego.Run(":8089")
}
