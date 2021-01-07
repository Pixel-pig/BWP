package main

import (
	"BWPjy/dbMysql"
	_ "BWPjy/routers"
	"github.com/astaxie/beego"
)

func main() {
	dbMysql.ConnectDB()
	beego.Run()
}

