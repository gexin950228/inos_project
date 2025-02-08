package routers

import (
	"github.com/astaxie/beego"
	"inos_project/controllers"
	"inos_project/controllers/login"
	"inos_project/controllers/user"
)

func init() {
	// 未登录也可以访问的url
	beego.Router("/", &login.LoginController{})
	beego.Router("/change_captcha", &login.LoginController{}, "get:ChangeCaptcha")

	// 登陆后才能访问的url
	beego.Router("/inner/home", &controllers.HomeController{})
	beego.Router("/inner/welcome", &controllers.HomeController{}, "get:Welcome")
	beego.Router("/inner/user-list", &user.ExecUserController{})
	beego.Router("/inner/user-add", &user.ExecUserController{}, "*:AddUser")
}
