package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int       `json:"id" orm:"pk;auto"`
	UserName   string    `form:"username" json:"user_name" orm:"unique;size(50);column(user_name);description(用户名)"`
	Password   string    `form:"password" json:"password" orm:"column(password);description(密码)"`
	Age        int       `form:"age" json:"age" orm:"size(10);column(age);description(年龄);null"`
	Gender     int       `form:"gender" json:"gender" orm:"size(10);column(gender);description(性别);null"`
	Phone      string    `form:"phone" json:"phone" orm:"size(50);description(电话号码);column(phone);description(电话号码)"`
	Addr       string    `form:"addr" json:"addr" orm:"size(50);column(addr);description(住址);null"`
	CreateTime time.Time `orm:"type(datetime);auto_now;column(create_time);null"`
	IsDeleted  int       `orm:"type(int);column(is_deleted);description(账号是否可以)" json:"is_deleted"`
	Email      string    `orm:"column(email);size(100);description(邮箱)" json:"email" form:"email"`
	IsActice   int       `orm:"column(is_active);type(int)" form:"is_ctive" json:"is_ctive"`
}

func (u *User) TableName() string {
	return "user"
}

func init() {
	username := beego.AppConfig.String("username")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	password := beego.AppConfig.String("password")
	database := beego.AppConfig.String("database")
	dst := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local", username, password, host, port, database)
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = orm.RegisterDataBase("default", "mysql", dst)
	if err != nil {
		fmt.Println(err.Error())
	}
	orm.RegisterModel(new(User))
}
