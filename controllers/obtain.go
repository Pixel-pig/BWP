package controllers

import "github.com/astaxie/beego"

type Obtain struct {
	beego.Controller
}

func (o *Obtain)Get()  {
	o.TplName="obtain.html"
}