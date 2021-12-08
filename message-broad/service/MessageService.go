package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-broad/Dao"
	"message-broad/Struct"
)

// FindAllPMsg 查找所有留言或评论
func FindAllPMsg(c *gin.Context, db *sql.DB, pid int) error {
	err := Dao.FindAll(c, db, pid)
	return err
}

// SendMsg 添加留言
func SendMsg(db *sql.DB, username string, msg string, pid int) (int, error) {
	id, err := Dao.InsertMsg(db, username, msg, pid)
	return id, err
}

// FindAllComments 展示评论
func FindAllComments(c *gin.Context, db *sql.DB, pid int) error {
	err := Dao.ListComments(c, db, pid)
	return err
}

// GetOneMsg 查找留言或评论
func GetOneMsg(db *sql.DB, id string) (Struct.Info, bool) {
	u, flag := Dao.FindTheMsg(db, id)
	return u, flag
}

// DeleteMsg 删除留言
func DeleteMsg(db *sql.DB, id string) {
	Dao.DeleteMsg(db, id)
}

func ThumbAdd(db *sql.DB, id string, username string) {
	Dao.ThumbAdd(db, id, username)
}
