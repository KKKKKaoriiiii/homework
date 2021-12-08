package tool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// GetDb 获取数据库
func GetDb() *sql.DB {
	return Db
}

// OpenDb 连接数据库
func OpenDb() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	CheckErr(err)
	Db = db
	return
}
