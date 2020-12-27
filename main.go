package main

import (
	"BtcProject/db_mysql"
	_ "BtcProject/routers"
	"github.com/astaxie/beego"
)

func main() {

	db_mysql.ConnectSql()

	beego.Run()
}

