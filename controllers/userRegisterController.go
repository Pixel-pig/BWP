package controllers

import (
	"BWP/modles"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

func (r *RegisterController) Post() {
	//解析请求数据
	var user modles.User1
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("数据解析失败！请重试")
		return
	}
	//将解析数据保存到数据库
	_, err = user.SaveUser()
	if err != nil {
		r.Ctx.WriteString("数据保存失败！请重试")
		return
	}
	//注册成功，跳转页面
	r.TplName = "login.html"
}