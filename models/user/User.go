package user

import "BWP/db_mysql"

type User struct {
	Name     string `form:"username"`
	Password string `form:"pwd"`
}

//将数据存入数据库中
func (u *User) SaveUserInfo() (int64,error) {
	result, err := db_mysql.DB.Exec("insert into user "+
		"(username, pwd)"+
		"values (?, ?)", u.Name, u.Password)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
