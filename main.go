package main

import (
	"BTCproject/BTC_mysql"
	_ "BTCproject/routers"
	"BTCproject/util"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	BTC_mysql.Btc()
	beego.SetStaticPath("js","./static/js")
	beego.SetStaticPath("css","./static/css")
	beego.SetStaticPath("img","./static/img")
	hash:=util.GetBestBlockHash()
	fmt.Println(hash)
	block:=util.GetBlock(hash)
	fmt.Println(block)

	beego.Run()
}

