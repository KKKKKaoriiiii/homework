package controller

import (
	"github.com/gin-gonic/gin"
	"message-broad/service"
	"message-broad/tool"
	"strconv"
)

// MessageRouter 注册路由
func MessageRouter(engine *gin.Engine) {
	engine.POST("/msg", SendMsg)
	engine.POST("/anonymousMsg", anonymousMsg)
	engine.POST("/msgDelete", AdminMiddleWare, deleteMsg)
	engine.GET("/msgList", listMsg)
	engine.GET("/msg", getOneMsg)
	engine.POST("/msg/:id/comment", SendComment)
	engine.POST("/msg/:id/messages", listComment)
	engine.POST("/thumb", thumb)
}

// thumb 点赞
func thumb(c *gin.Context) {
	id := c.PostForm("id")
	db := tool.GetDb()
	username := tool.CheckLog(c)
	if username == "" {
		tool.PrintInfo(c, "你还没有登录！")
		return
	}
	service.ThumbAdd(db, id, username)
}

// SendMsg 发送留言
func SendMsg(c *gin.Context) {
	msg := c.PostForm("msg")
	username := tool.CheckLog(c)
	if username == "" {
		tool.PrintInfo(c, "你还没有登录！")
		return
	}
	db := tool.GetDb()
	id, err2 := service.SendMsg(db, username, msg, 0)
	tool.CheckErr(err2)
	tool.PrintInfo(c, "你已经留言成功！"+"id为"+strconv.Itoa(id))
	return
}

// anonymousMsg 发送匿名留言
func anonymousMsg(c *gin.Context) {
	msg := c.PostForm("msg")
	username := "anonymousUser"
	db := tool.GetDb()
	id, err2 := service.SendMsg(db, username, msg, 0)
	tool.CheckErr(err2)
	tool.PrintInfo(c, "你已经留言成功！"+strconv.Itoa(id))
	return
}

// getOneMsg 获取一条留言
func getOneMsg(c *gin.Context) {
	id := c.Param("id")
	db := tool.GetDb()
	u, flag := service.GetOneMsg(db, id)
	if flag == false {
		tool.PrintInfo(c, "无该id对于的留言。")
	} else {
		tool.PrintMsg(c,u)
	}
	return
}

// listMsg 获取全部留言
func listMsg(c *gin.Context) {
	db := tool.GetDb()
	err := service.FindAllPMsg(c, db, 0)
	if err != nil {
		tool.PrintInfo(c, "No message!")
		return
	}
	return
}

// SendComment 发送评论
func SendComment(c *gin.Context) {
	msg := c.PostForm("msg")
	username := tool.CheckLog(c)
	if username == "" {
		tool.PrintInfo(c, "你还没有登录！")
		return
	}
	pid, err := strconv.Atoi(c.PostForm("toWhich"))
	if err != nil {
		tool.PrintInfo(c, "你输入的评论地址有误！")
		return
	}
	db := tool.GetDb()
	_, err2 := service.SendMsg(db, username, msg, pid)
	tool.CheckErr(err2)
	return
}

// listComment 显示所有评论
func listComment(c *gin.Context) {
	db := tool.GetDb()
	pid := c.PostForm("toWhich")
	pidNum, _ := strconv.Atoi(pid)
	err := service.FindAllComments(c, db, pidNum)
	if err != nil {
		tool.PrintInfo(c, "No message!")
		return
	}
	return
}

// deleteMsg 删除留言
func deleteMsg(c *gin.Context) {
	id := c.PostForm("id")
	db := tool.GetDb()
	service.DeleteMsg(db, id)
}
