package controllers

import (
	"BTCproject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type Index struct {
	beego.Controller
}

func (i *Index) Get() {
	i.TplName="index.html"
}
func (i *Index) Post() {
	var users models.Users
	err := i.ParseForm(&users)
	if err != nil {
		i.Ctx.WriteString("抱歉，解析数据错误，请重试！")
		return
	}
	_, err = users.SaveUser()
	if err != nil {
		fmt.Println(err)
		i.Ctx.WriteString("抱歉，用户注册失败，请重试！")
		return
	}
	i.TplName = "index.html"
}
