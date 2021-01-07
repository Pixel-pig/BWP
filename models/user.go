package models

import (
	"BWPjy/dbMysql"
)

type User struct {
	Id int `form:"id"`
	Name string `form:"name"`
	Password string `form:"password"`
}

func (u User) Queryuser () (*User , error) {
	row := dbMysql.DB.QueryRow("select * from user where name = ? and password = ?",u.Name,u.Password)
	err := row.Scan(&u.Id,&u.Name,&u.Password)
	if err != nil {
		return nil , err
	}
	return &u , nil
}
func (u User) Adduser () (int64,error) {
	result , err := dbMysql.DB.Exec("insert into user (name , password) values (?,?)",u.Name,u.Password)
	if err != nil {
		return -1 , err
	}
	id , err := result.RowsAffected()
	if err != nil {
		return -1 , err
	}
	return id , nil
}