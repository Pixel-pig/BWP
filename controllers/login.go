package controllers

import (
	"BTCproject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type Login struct {
	beego.Controller
}

func (l *Login) Post() {
	var users models.Users
	err := l.ParseForm(&users)
	if err != nil {
		fmt.Println(err)
		l.Ctx.WriteString("抱歉，用户解析失败,请重试!")
		return
	}
	u, err := users.QueryUser()
	if err != nil {
		fmt.Println(err)
		l.Ctx.WriteString("抱歉，用户登入失败，请重试！")
		return
	}
	fmt.Println(u)
	//l.Data["Phone"]=u.Phone
	l.TplName = "Btc_home.html"
}
