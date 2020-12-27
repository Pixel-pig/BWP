package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectSql()  {
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")
	dbCharset := config.String("db_charset")

	//连接数据库
	connUrl := dbUser + ":" + dbPassword + "@tcp" + dbIP + dbName + dbCharset
	DB, err := sql.Open(dbDriver,connUrl)
	if err != nil {
		panic("数据库连接失败！请检查配置")
	}
	fmt.Println("数据库连接成功！")
	Db = DB
}
