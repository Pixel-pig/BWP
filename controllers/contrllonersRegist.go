package controllers

import (
	"BWPjy/models"
	"github.com/astaxie/beego"
)

type RegistControllers struct {
	beego.Controller
}

func (r *RegistControllers)Get() {
	r.TplName = "RegistPage.html"
}

func (r *RegistControllers)Post() {
	var user models.User

	err := r.ParseForm(&user)
	if err != nil {
		r.TplName = "ErrorPage.html"
	}
	_ , err = user.Adduser()
	if err != nil {
		r.TplName = "ErrorPage.html"
	}
	r.TplName = "LoginPage.html"
}