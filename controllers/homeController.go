package controllers

import (
	"BWP/modles"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (h *HomeController) Get() {
	h.TplName = "home.html"
}

func (h *HomeController) Post() {
	var result modles.RPCRequest
	err := h.ParseForm(&result)
	if err != nil {
		h.Ctx.WriteString("数据解析失败！")
	}
	h.TplName = ""
}
