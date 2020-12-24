package controllers

import "github.com/astaxie/beego"

type Btc_home struct {
	beego.Controller
}

func (b *Btc_home)Get()  {
	b.TplName="Btc_home.html"
}
