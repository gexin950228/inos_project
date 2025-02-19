package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"inos_project/models/user"
	"inos_project/utils"
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
	var insertUser user.User
	err := u.ParseForm(&insertUser)
	if err != nil {
		logs.Error(fmt.Sprintf("注册用户参数绑定失败，错误信息:%s", err.Error()))
	} else {
		fmt.Printf("user:%#v\n", insertUser)
	}

	// 初始化相应结构体
	repMsg := make(map[string]string)

	o := orm.NewOrm()
	qs := o.QueryTable("user")
	err = qs.Filter("username", insertUser.UserName).One(&insertUser)
	if err != nil {
		logs.Error(fmt.Sprintf("查询数据库中名字为:%s，邮箱为%s的用户失败，错误信息为: %s", insertUser.UserName, insertUser.Email, err.Error()))
	}
	if insertUser.Id > 0 {
		repMsg["code"] = "0"
		repMsg["msg"] = fmt.Sprintf("查询数据库中名字为:%s，邮箱为%s的用户已存在", insertUser.UserName, insertUser.Email)
	} else {
		md5Pass := utils.EncryptPassword(insertUser.Password)
		insertUser.Password = md5Pass
		insertId, err := o.Insert(&insertUser)
		if err != nil {
			logs.Error(err.Error())
			repMsg["code"] = "1"
			repMsg["msg"] = fmt.Sprintf("数据库插入失败，错误信息:%s", err.Error())
		} else {
			repMsg["code"] = "0"
			repMsg["msg"] = fmt.Sprintf("用户插入成功，用户id是: %v", insertId)
		}
	}
	u.ServeJSON()
}

func (u *ExecUserController) ExecModifyUser() {
	u.ServeJSON()
}
