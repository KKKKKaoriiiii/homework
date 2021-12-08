package tool

import (
	"github.com/gin-gonic/gin"
	"message-broad/Struct"
	"strconv"
)

// PrintInfo 输出字符串
func PrintInfo(c *gin.Context, str string) {
	c.String(200, str+"\n")
}

// PrintMsg 输出留言或评论
func PrintMsg(c *gin.Context, u Struct.Info) {
	if u.Thumb == 0 {
		if u.Pid == 0 {
			c.String(200, "\nid"+strconv.Itoa(u.Id)+":"+"用户名为"+u.Username+"留言内容:\n"+u.Msg+"\n评论时间为"+strconv.FormatInt(u.Time, 10)+"\n评论数为"+strconv.Itoa(u.CommentNum))
		} else {
			c.String(200, "\nid"+strconv.Itoa(u.Id)+":"+"用户名为"+u.Username+"对ID为"+strconv.Itoa(u.Pid)+"的评论内容:\n"+u.Msg+"\n评论时间为"+strconv.FormatInt(u.Time, 10)+"\n评论数为"+strconv.Itoa(u.CommentNum))
		}
	} else {
		if u.Pid == 0 {
			c.String(200, "\nid"+strconv.Itoa(u.Id)+":"+"用户名为"+u.Username+"留言内容:\n"+u.Msg+"\n评论时间为"+strconv.FormatInt(u.Time, 10)+"\n评论数为"+strconv.Itoa(u.CommentNum)+
				"\n点赞数为"+strconv.Itoa(u.Thumb)+u.Liker+"为该留言点了赞！\n")
		} else {
			c.String(200, "\nid"+strconv.Itoa(u.Id)+":"+"用户名为"+u.Username+"对ID为"+strconv.Itoa(u.Pid)+"的评论内容:\n"+u.Msg+"\n评论时间为"+strconv.FormatInt(u.Time, 10)+"\n评论数为"+strconv.Itoa(u.CommentNum)+
				"\n点赞数为"+strconv.Itoa(u.Thumb)+u.Liker+"为该评论点了赞！\n")
		}
	}
}
