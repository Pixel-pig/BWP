package models

import (
	"BTCproject/BTC_mysql"
	"BTCproject/util"
)

type Users struct {
	Id         int64       `form:"id"`
	Phone      int64       `form:"phone"`
	User       string      `form:"user"`
	Password   string      `form:"password"`
	Method     string      `form:"method"`
	Parameter  interface{} `form:"parameter"`
	Parameter1 interface{} `form:"parameter1"`
	Parameter2 interface{} `form:"parameter2"`
}

func (u Users) SaveUser() (int64, error) {
	u.Password = util.MD5HashString(u.Password)
	row, err := BTC_mysql.Db.Exec("insert into userindex(phone,password,user)"+
		"values(?,?,?)", u.Phone, u.Password, u.User)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (u Users) QueryUser() (*Users, error) {
	u.Password = util.MD5HashString(u.Password)
	row := BTC_mysql.Db.QueryRow("select id, phone, user from userindex where phone = ? and password = ?",
		u.Phone, u.Password)
	err := row.Scan(&u.Id, &u.Phone, &u.User)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
