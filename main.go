package main

import (
	"BTCproject/BTC_mysql"
	_ "BTCproject/routers"
	"github.com/astaxie/beego"
)

func main() {
	BTC_mysql.Btc()
	beego.SetStaticPath("js","./static/js")
	beego.SetStaticPath("css","./static/css")
	beego.SetStaticPath("img","./static/img")
	beego.Run()
}

