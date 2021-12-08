package Dao

import (
	"database/sql"
	"fmt"
	"message-broad/Struct"
	"message-broad/tool"
)

// InsertUser 插入一条信息
func InsertUser(db *sql.DB, username string, password string, answer string, question string) error {
	sqlStr := "insert into user(username, password, secretQuestion, secretAnswer) values (?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, username, password, question, answer)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	} else {
		return nil
	}
}

// UpdateRowDemo 修改密码
func UpdateRowDemo(db *sql.DB, pwd string, username string) error {
	sqlStr := "update user set password=? where username = ?"
	_, err := db.Exec(sqlStr, pwd, username)
	if err != nil {
		fmt.Println("update failed, err:", err)
		return err
	} else {
		return nil
	}
}

// FindUser 查找用户
func FindUser(db *sql.DB, username string) (Struct.User, bool) {
	sqlStr := "select username, password, secretQuestion, secretAnswer from user where username = ?"
	var u Struct.User
	err := db.QueryRow(sqlStr, username).Scan(&u.Username, &u.Pwd, &u.Question, &u.Answer)
	if err != nil {
		fmt.Println("scan failed err:", err)
		return u, false
	}
	return u, true
}

// AddSecret 添加密保
func AddSecret(db *sql.DB, question string, answer string, username string) {
	sqlStr := "update user set secretQuestion = ?, secretAnswer = ? where username = ?"
	_, err := db.Exec(sqlStr, question, answer, username)
	tool.CheckErr(err)
}
