package main

import (
	"BWP/db_mysql"
	_"BWP/routers"
	"github.com/astaxie/beego"
)

func main() {

	db_mysql.ConnectSql()

	beego.Run()
}

