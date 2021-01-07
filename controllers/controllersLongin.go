package controllers

import (
	"BWPjy/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LonginControllers struct {
	beego.Controller
}

func (l *LonginControllers) Get() {
	l.TplName = "LoginPage.html"
}

func (l *LonginControllers) Post() {
	var user models.User

	err := l.ParseForm(&user)
	if err != nil {
		l.TplName = "ErrorPage.html"
		return
	}
	fmt.Println(user.Name,user.Password)
	 _ , err =user.Queryuser()
	if err != nil {
		l.TplName = "ErrorPage.html"
		return
	}
	l.TplName = "home.html"
}