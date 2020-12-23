package main

import (
	_ "BWP/routers"
	"github.com/astaxie/beego"
)

func main() {
	//加载静态资源
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/js","./static/js")

	beego.Run()
}

