package controllers

import (
	"BWP/modles"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}

func (l *LoginController) Post() {
	var user modles.User1
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("数据解析失败！请重试")
		return
	}
	u,err := user.QueryUser()
	if err != nil {
		l.Ctx.WriteString("数据库查询失败！请重试")
	}
	l.Data["Name"] = u.Name
	l.TplName = "home.html"
}
