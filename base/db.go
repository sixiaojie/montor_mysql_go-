package base

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func JDBC(host string)(db *sql.DB,err error){
	user := cdbinfo[env]["user"]
	password := cdbinfo[env]["password"]
	port := "3306"
	database := "information_schema"
	charset := "utf8"
	connect_S := user+":"+password+"@tcp"+"("+host+":"+port+")"+"/"+database+"?"+"charset="+charset
	db,err = sql.Open("mysql",connect_S)
	if err != nil {
		Logger.Warning("db:%s 无法连接",host)
		return nil,err
	}
	return db,nil
}

