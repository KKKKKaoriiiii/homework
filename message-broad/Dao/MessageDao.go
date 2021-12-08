package Dao

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-broad/Struct"
	"message-broad/tool"
	"strconv"
	"time"
)

// InsertMsg 添加留言
func InsertMsg(db *sql.DB, username string, msg string, pid int) (int, error) {
	sqlStr := "insert into Info(username, msg, pid, time) values (?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, username, msg, pid, time.Now().Unix())
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0, err
	} else {
		if pid != 0 {
			commentAdd(db, pid)
		}
		var id int
		err1 := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
		tool.CheckErr(err1)
		return id, nil
	}
}

// FindTheMsg 查找留言或评论
func FindTheMsg(db *sql.DB, id string) (Struct.Info, bool) {
	sqlStr := "select username, msg, id, pid, time, commentNum, thumbs_up, liker from Info where id = ?"
	row := db.QueryRow(sqlStr, id)
	var u Struct.Info
	err := row.Scan(&u.Username, &u.Msg, &u.Id, &u.Pid, &u.Time, &u.CommentNum, &u.Thumb, &u.Liker)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return u, false
	} else {
		return u, true
	}
}

// CloseDb 关闭数据库
func CloseDb(row *sql.Rows) {
	err := row.Close()
	tool.CheckErr(err)
}

// commentAdd 添加评论数
func commentAdd(db *sql.DB, pid int) {
	var commentNum int
	sqlStr := "select commentNum from Info where id = ?"
	row := db.QueryRow(sqlStr, pid)
	err := row.Scan(&commentNum)
	tool.CheckErr(err)
	commentNum++
	sqlStr2 := "update info set commentNum = ? where id = ?"
	_, err2 := db.Exec(sqlStr2, commentNum, pid)
	tool.CheckErr(err2)
}

// DeleteMsg 删除留言
func DeleteMsg(db *sql.DB, id string) {
	sqlStr := "delete from info where id = ?"
	idNum, _ := strconv.Atoi(id)
	ret, err := db.Exec(sqlStr, idNum)
	tool.CheckErr(err)
	n, err := ret.RowsAffected()
	tool.CheckErr(err)
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// FindAll 查找所有留言或评论
func FindAll(c *gin.Context, db *sql.DB, pid int) error {
	sqlStr := "select username, msg, id, time, commentNum,thumbs_up,liker from Info where pid = ?"
	rows, err := db.Query(sqlStr, pid)
	defer CloseDb(rows)
	var u Struct.Info
	for rows.Next() {
		err = rows.Scan(&u.Username, &u.Msg, &u.Id, &u.Time, &u.CommentNum, &u.Thumb, &u.Liker)
		u.Pid = pid
		tool.PrintMsg(c, u)
	}
	return err
}

// ListComments 展示评论
func ListComments(c *gin.Context, db *sql.DB, pid int) error {
	var u Struct.Info
	sqlStr2 := "select username, msg, id, time, commentNum,thumbs_up,liker from Info where pid = ?"
	rows, err2 := db.Query(sqlStr2, pid)
	defer CloseDb(rows)
	if err2 != nil {
		return err2
	}
	for rows.Next() {
		err := rows.Scan(&u.Username, &u.Msg, &u.Id, &u.Time, &u.CommentNum, &u.Thumb, &u.Liker)
		u.Pid = pid
		if err != nil {
			fmt.Println(err)
			return err
		}
		tool.PrintMsg(c, u)
		var commentNum1 int
		sqlStr := "select commentNum from Info where id = ?"
		err4 := db.QueryRow(sqlStr, u.Id).Scan(&commentNum1)
		if err4 != nil {
			return err4
		}
		if commentNum1 != 0 {
			err3 := ListComments(c, db, u.Id)
			if err3 != nil {
				return err3
			}
		}
	}
	return nil
}

func ThumbAdd(db *sql.DB, id string, username string) {
	idNum, _ := strconv.Atoi(id)
	sqlStr := "select thumbs_up, liker from info where id = ?"
	row := db.QueryRow(sqlStr, idNum)
	var thump int
	var liker string
	err := row.Scan(&thump, &liker)
	tool.CheckErr(err)
	thump++
	liker = liker + "," + username
	sqlStr = "update info set thumbs_up = ?, liker = ? where id = ?"
	_, err = db.Exec(sqlStr, thump, liker, id)
	tool.CheckErr(err)
}
