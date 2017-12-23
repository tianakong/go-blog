package controllers

import (
	"blog/models"
	"time"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	BaseController
} 

func (this *UserController) Register() {
	username := this.GetString("email")
	password := this.GetString("password")
	if username != "" {
		_,err := models.UserGetByName(username)
		if err == nil {
			this.ajaxReturn(-1, "用户名已存在")
		}
		user := new(models.User)
		user.Username = username
		user.Password = this.md5(password)
		_,err = models.UserAdd(user)
		if err == nil {
			this.ajaxReturn(1, "注册成功")
		} else {
			this.ajaxReturn(-2, "注册失败")
		}
	}

	this.TplName = "reg.tpl"
}

func (this *UserController) Login() {
	userId := this.GetSession("userId")
	if userId != nil {
		this.Redirect("/", 302)
	}
	username := this.GetString("email")
	password := this.GetString("password")
	if username != "" {
		user,err := models.UserGetByName(username)
		if err != nil {
			this.ajaxReturn(-1, "用户不存在")
		}
		if user.Password != this.md5(password) {
			this.ajaxReturn(-2, "密码错误")
		}
		this.SetSession("userId", user.Id)
		this.SetSession("userName", user.Username)
		user.Logintime = time.Now()
		orm.NewOrm().Update(user, "Logintime")
		this.ajaxReturn(1, "登录成功")
	}

	this.TplName = "login.tpl"
}

func (this *UserController) Loginout() {
	this.DelSession("userId")
	this.DelSession("userName")
	this.Redirect("/", 302)
}