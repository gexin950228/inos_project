package routers

import (
	"github.com/astaxie/beego"
	"inos_project/controllers/login"
)

func init() {
	beego.Router("/", &login.LoginController{})
	beego.Router("/change_captcha", &login.LoginController{}, "get:ChangeCaptcha")
}
