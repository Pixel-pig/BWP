package controllers

import "github.com/astaxie/beego"

type HomeControllers struct {
	beego.Controller
}

func (h *HomeControllers) Get()  {
	h.TplName = "home.html"
}

