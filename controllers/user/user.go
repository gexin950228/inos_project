package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"inos_project/models/user"
	"math"
)

type ExecUserController struct {
	beego.Controller
}

func (u *ExecUserController) Get() {
	pageId, err := u.GetInt("pageId")
	var currentPage int
	if err != nil {
		logs.Error(fmt.Sprintf("获取pageId出错，错误信息为: %s", err.Error()))
		currentPage = 1
	} else {
		if pageId > 0 {
			currentPage = pageId
		} else {
			currentPage = 1
		}
	}
	limitNum := 3
	offsetNum := limitNum * (currentPage - 1)
	orm.Debug = true
	o := orm.NewOrm()
	var users []user.User
	qs := o.QueryTable("user")
	totalUserNum, _ := qs.Count()
	totalPageNum := int(math.Ceil(float64(totalUserNum) / float64(limitNum)))
	_, err = qs.Limit(limitNum).Offset(offsetNum).All(&users)
	if err != nil {
		logs.Error(fmt.Sprintf("查询用户数据报错，错误信息为:%s", err.Error()))
		return
	}
	u.Data["currentPage"] = currentPage
	var prePage int
	if currentPage > 1 {
		prePage = currentPage - 1
	} else {
		prePage = 0
	}
	var nextPage int
	if currentPage < totalPageNum {
		nextPage = currentPage + 1
	} else {
		nextPage = totalPageNum
	}
	u.Data["prePage"] = prePage
	u.Data["nextPage"] = nextPage
	u.Data["users"] = users
	u.Data["totalUserNum"] = totalUserNum
	u.Data["totalPageNum"] = totalPageNum
	u.Data["currentPage"] = currentPage
	u.TplName = "user/member-list.html"
}

func (u *ExecUserController) AddUser() {
	u.TplName = "user/member-add.html"
}

func (u *ExecUserController) EditUser() {
	u.TplName = "user/member-edit.html"
}

func (u *ExecUserController) ExecAddUser() {
	u.ServeJSON()
}
