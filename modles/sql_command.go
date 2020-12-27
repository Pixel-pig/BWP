package modles

import (
	db_mysql2 "BtcProject/db_mysql"
	"fmt"
)

type User1 struct {
	Id       int    `form:"id"`
	Name     string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

func (u User1) SaveUser()(int64,error) {
	result,err := db_mysql2.Db.Exec("insert into login(phone, password,name)"+"values(?,?,?)",u.Phone,u.Password,u.Name)
	if err != nil {
		return -1,err
	}
	id,err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}

func (u User1) QueryUser()(*User1,error) {
	row := db_mysql2.Db.QueryRow("select name from login where name = ? and password = ?",u.Name,u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return &u,err
}