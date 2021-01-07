package dbMysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	Config := beego.AppConfig
	dbdriver := Config.String("db_driverName")
	dbuser := Config.String("db_user")
	dbpwd := Config.String("db_password")
	dbip := Config.String("db_ip")
	dbname := Config.String("db_name")

	dburl := dbuser + ":" + dbpwd + "@tcp(" + dbip + ")/" + dbname + "?charset=utf8"

	Db , err := sql.Open(dbdriver,dburl)
	if err != nil {
		panic("数据库连接失败")
	}

	DB = Db
}
