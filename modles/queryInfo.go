package modles

import (
	"BWP/db_mysql"
)

type QueryInfo struct {
	Method string `form:"method"`
	Param  string `form:"param"`
	Param1 string `form:"param1"`
	Param2 string `form:"param2"`
}

func (q QueryInfo) SeverBlock()(int64,error) {
	result,err := db_mysql.Db.Exec("insert into block(method,param,param1,param2)"+"values(?,?,?,?)",q.Method,q.Param,q.Param1,q.Param2)
	if err != nil {
		return -1,err
	}
	id,err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}
