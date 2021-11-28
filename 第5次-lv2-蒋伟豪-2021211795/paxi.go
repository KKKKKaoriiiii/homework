package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"strings"
)
var r=gin.Default()
var i = 0
func main(){
	Init()
	menu()
}
func menu() {
	fmt.Println("欢迎来到世界树")
	fmt.Println("请发送get到http://127.0.0.1:8080/login以登录")
	fmt.Println("请向http://127.0.0.1:8080/reg发送post以注册")
	fmt.Println("请向http://127.0.0.1:8080/changePassword发送fet以修改密码")
	var ok1=false
	var ok2=false
	r.GET("/login",auth,func(c *gin.Context) {
		input1:=c.Query("username")
		input2:=c.Query("password")
		for t:=0;t<i;t++{
			m1, _ :=c.Get(strconv.Itoa(t)+"1")
			m2, _ :=c.Get(strconv.Itoa(t)+"2")
			if m1 == input1 {
				ok1=true
			}
			if m2 == input2{
				ok2=true
			}
			if ok1 && ok2{break}
		}
		if !ok1 {
			c.String(403,"账号不存在")
			return
		}
		if !ok2 {
			c.String(403,"密码错误")
			return
		}
		cookie.Name = input1
		cookie.Password = input2
		c.String(200,"登录成功")
	})
	r.POST("/reg",auth,func(c *gin.Context) {
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		name := username
		for t:=0;t<i;t++{
			m1, _ := c.Get(strconv.Itoa(t)+"1")
			ok := m1 == username
			if ok {
				c.String(403,"账号已存在")
				return
			}
		}
		if name == "" || checkIfSensitive(name) {
			c.String(403,"用户名不合法")
			return
		}

		if !checkPasswordLegal(password) {
			c.String(403,"密码不合法")
			return
		}
		password = defaultEncrypt(password)
		cookie.Name = username
		cookie.Password = password
		c.SetCookie(strconv.Itoa(i)+"1",username,3600,"/","",false,true)
		c.SetCookie(strconv.Itoa(i)+"2",password,3600,"/","",false,true)
		c.String(200,"注册成功，已自动登录")
	})
	r.GET("/changePassword",auth, func(c *gin.Context) {
		if cookie.Name == "" {
			c.String(200,"你还没有登录")
			return
		}
		password := c.Query("password")
		if !checkPasswordLegal(password) {
			c.String(200,"密码不合法")
			return
		}
		password = defaultEncrypt(password)
		cookie.Password = password
		for t:=0;t<i;t++{
			m1, _ :=c.Get(strconv.Itoa(t)+"1")
			if m1 == cookie.Name {
				c.SetCookie(strconv.Itoa(t)+"2",password,3600,"/","",false,true)
				break
			}
		}
		c.String(200,"修改密码成功！")
	})
	r.Run()
}

var (
	cookie         UserInfo
	sensitiveWords = make([]string, 0)
)
func checkIfSensitive(s string) bool {
	for _, word := range sensitiveWords {
		if strings.Contains(s, word) {
			return true
		}
	}
	return false
}
func checkPasswordLegal(password string) bool {
	return len(password) > 6
}
func defaultEncrypt(raw string) string {
	return encrypt(raw, encryptSalt)
}
func encrypt(raw string, salt string) string {
	has := md5.New()
	_, err := io.WriteString(has, raw)
	if err != nil {
		fmt.Println(err)
	}
	tem := has.Sum([]byte(salt))
	Result := hex.EncodeToString(tem)
	return Result
}

const (
	encryptSalt = "BaiLanDeShen"
)
type UserInfo struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *UserInfo) ifPasswordCorrect(password string) bool {
	return defaultEncrypt(password) == u.Password
}

func Init() {
	sensitiveWords = append(sensitiveWords, "你妈", "傻逼")
}
func auth(c *gin.Context){
	for {
		username, err := c.Cookie(strconv.Itoa(i)+"1")
		if err != nil{
			return
		}
		c.Set(strconv.Itoa(i)+"1", username)
		password, _ := c.Cookie(strconv.Itoa(i)+"2")
		c.Set(strconv.Itoa(i)+"1", password)
		i++
	}
}
