package controller

import (
	"github.com/gin-gonic/gin"
	"message-broad/tool"
)

// AdminRouter 注册路由
func AdminRouter(engine *gin.Engine) {
	engine.GET("/msgClear", AdminMiddleWare, MsgClear)
	engine.POST("/userDelete", AdminMiddleWare, userDelete)
}

// AdminMiddleWare 管理员权限确认
func AdminMiddleWare(c *gin.Context) {
	username := tool.CheckLog(c)
	if username == "cjw" {
		return
	}
	tool.PrintInfo(c, "你不是管理员！")
	c.Abort()
	return
}

// MsgClear 清空所有留言
func MsgClear(c *gin.Context) {
	db := tool.GetDb()
	sqlStr := "truncate table info"
	_, err := db.Exec(sqlStr)
	tool.CheckErr(err)
	tool.PrintInfo(c, "清空成功！")
}

// userDelete 删除用户
func userDelete(c *gin.Context) {
	username := c.PostForm("username")
	db := tool.GetDb()
	sqlStr := "delete from user where username = ?"
	_, err := db.Exec(sqlStr, username)
	tool.CheckErr(err)
	tool.PrintInfo(c, "注销成功！")
}
