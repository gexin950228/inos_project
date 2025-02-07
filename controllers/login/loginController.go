package login

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"inos_project/models/user"
	"inos_project/utils"
)

type LoginController struct {
	beego.Controller
}

var store = base64Captcha.DefaultMemStore

type Captcha struct {
	Code    int
	Id      string `json:"id"`
	Captcha string
}

func GetCaptcha() (id, base64 string, err error) {
	// 生产driver, 参数:height int, width int, noiseCount int, showLineOptions int, bgColor *color.RGBA, fonts
	rgbColor := color.RGBA{}
	fonts := []string{"wqy-microhei.ttc"}
	driver := base64Captcha.NewDriverMath(50, 140, 2, 2, &rgbColor, fonts)

	// 使用store和driver生成验证码实例
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 生成验证码
	id, base64, err = captcha.Generate()
	return id, base64, err
}

func (c *LoginController) Get() {
	id, base64, err := GetCaptcha()
	if err != nil {
		fmt.Println(err.Error())
		logs.Error(err)
		return
	}
	c.Data["captcha"] = Captcha{
		Id:      id,
		Captcha: base64,
	}
	c.TplName = "login/login.html"
}

func (c *LoginController) ChangeCaptcha() {
	msg := map[string]string{}
	id, base64, err := GetCaptcha()
	if err != nil {
		logs.Error(err)
		msg["err"] = fmt.Sprintf("生成验证码失败,错误信息: %s", err.Error())
		c.Data["json"] = Captcha{Id: "", Code: 0, Captcha: ""}
	} else {
		c.Data["json"] = Captcha{Code: 1, Id: id, Captcha: base64}
	}
	c.ServeJSON()
}

type LoginInfo struct {
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
	CaptchaId string `form:"captchaId" json:"captchaId"`
	Captcha   string `form:"captcha" json:"captcha"`
}

func VerifyCaptcha(id, base64 string) bool {
	return store.Verify(id, base64, true)
}

func (c *LoginController) Post() {
	var loginInfo LoginInfo
	err := c.ParseForm(&loginInfo)
	if err != nil {
		logs.Error(err)
		return
	}
	repResult := map[string]interface{}{}
	// 验证码校验
	verifyCodeCheck := VerifyCaptcha(loginInfo.CaptchaId, loginInfo.Captcha)
	//用户名密码校验
	if verifyCodeCheck {
		pbs := utils.EncryptPassword(loginInfo.Password)
		o := orm.NewOrm()
		exist := o.QueryTable("user").Filter("username", loginInfo.Username).Filter("password", pbs).Exist()
		var loginUser user.User
		if exist {
			o.QueryTable("user").Filter("username", loginInfo.Username).One(&loginUser)
			c.SetSession("userId", loginUser.Id)
			repResult["code"] = 0
			repResult["msg"] = "用户名密码正确，登陆成功"
		} else {
			repResult["code"] = 1
			repResult["msg"] = "用户名密码不正确"
		}
	} else {
		repResult["code"] = 1
		repResult["msg"] = "验证码不正确，请重新输入"
	}
	c.Data["json"] = repResult
	c.ServeJSON()
}
