package tool

import "github.com/gin-gonic/gin"

// CheckLog 检查是否登录
func CheckLog(c *gin.Context) string {
	cookie, err := c.Request.Cookie("login")
	if err == nil {
		return cookie.Value
	} else {
		return ""
	}
}
