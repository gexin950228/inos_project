package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"inos_project/models/user"
)

type ExecUserController struct {
	beego.Controller
}

func (u *ExecUserController) Get() {
	pageId, err := u.GetInt("page_id")
	if err != nil {
		pageId = 1
	}
	limitNum := 3
	offsetNum := limitNum * (pageId - 1)
	fmt.Println(limitNum, pageId)
	orm.Debug = true
	o := orm.NewOrm()
	var users []user.User
	qs := o.QueryTable("user")

	_, err = qs.Limit(limitNum).Offset(offsetNum).All(&users)
	if err != nil {
		logs.Error(fmt.Sprintf("查询用户数据报错，错误信息位:%s", err.Error()))
		return
	}
	u.Data["users"] = users
	u.TplName = "user/member-list.html"
}

func (u *ExecUserController) AddUser() {
	u.TplName = "user/member-add.html"
}
