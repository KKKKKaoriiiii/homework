package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var mine user

type user struct {
	username string
	password string
	secretQuestion string
	secretAnswer string
}
func main(){
	DB:=open()
	r:=gin.Default()
	login(r,DB)
	reg(r,DB)
	changeWord(r,DB)
	check(r,DB)
	whatQuest(r,DB)
	err:=r.Run()
	if err != nil {
		fmt.Println("捏麻麻的！")
	}
}

func whatQuest(r *gin.Engine, db *sql.DB) {
	r.POST("/whatQuest", func(c *gin.Context) {
		username:=c.PostForm("username")
		u,flag:=QueryRowDemoUsername(db,username)
		if flag && (u.secretQuestion != "") {
			c.String(200,"你的密保问题是%s",u.secretQuestion)
		}else if flag && (u.secretQuestion == ""){
			c.String(200,"你还没有设置密保！")
		}else {
			c.String(200,"无此账号！")
		}
	})
}

func changeWord(r *gin.Engine, db *sql.DB) {
	r.POST("/change",auth, func(c *gin.Context) {
		if mine.username == "" {
			c.String(403,"你还没有登录！")
			return
		}
		if mine.secretQuestion != "" {
			answer:=c.PostForm("answer")
			password:=c.PostForm("newPassword")
			if answer == mine.secretAnswer {
				updateRowDemo(db,password,mine.username)
				c.String(200,"修改完成！")
			}
		}else {
			password2:=c.PostForm("newPassword")
			updateRowDemo(db,password2,mine.username)
			c.String(200,"你还没有设置密保，修改完成!")
		}
	})
}
func reg(r *gin.Engine,db *sql.DB){
	var u user
	r.POST("/reg", func(c *gin.Context) {
		u.username = c.PostForm("username")
		u.password = c.PostForm("password")
		u.secretQuestion = c.DefaultPostForm("question","")
		u.secretAnswer = c.DefaultPostForm("answer","")
		_,flag := QueryRowDemoUsername(db,u.username)
		if flag {
			c.String(403,"该账号已经被注册!")
			return
		} else {
			insertRowDemo(db,u)
		}
	})
}
func QueryRowDemoUsername(db *sql.DB,username string) (user,bool) {
	sqlStr:="select username,password,secretQuestion,secretAnswer from user where username = ?"
	var u user
	err := db.QueryRow(sqlStr,username).Scan(&u.username,&u.password,&u.secretQuestion,&u.secretAnswer)
	if err != nil{
		fmt.Println("scan failed err:",err)
		return u,false
	}
	return u,true
}
func QueryRowDemoPassword(db *sql.DB,password string) (user,bool) {
	sqlStr := "select username,password,secretQuestion,secretAnswer from user where password = ?"
	var u user
	err := db.QueryRow(sqlStr,password).Scan(&u.username,&u.password,&u.secretQuestion,&u.secretAnswer)
	if err != nil{
		fmt.Println("scan failed err:",err)
		return u, false
	}
	return u, true
}
func open() *sql.DB {
	db,err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil{
		log.Fatal(err)
	}
	return db
	// 插入数据
}
func login(r *gin.Engine,db *sql.DB){
	r.POST("/login",auth, func(c *gin.Context) {
		value,err := c.Get("cookie")
		if err {
			mine1,flag := QueryRowDemoUsername(db, value.(string))
			if flag {
				mine = mine1
				c.String(200,"你已经自动登录，账号为%s",mine.username)
				return
			}
		}
		username := c.PostForm("username")
		password := c.PostForm("password")
		mine2,flag1 := QueryRowDemoUsername(db,username)
		_,flag2 := QueryRowDemoPassword(db,password)
		if flag2 && flag1 {
			mine = mine2
			c.SetCookie("gin_cookie",username,3600,"/","",false,true)
			c.String(200,"登陆成功!")
		}else {
			c.String(200,"账号或者密码错误！")
		}
	})
}
func auth(c *gin.Context){
	cookie,err:=c.Cookie("gin_cookie")
	if err != nil {
		return
	}
	c.Set("cookie",cookie)
}

func insertRowDemo(db *sql.DB,u user) {
	sqlStr := "insert into user(username, password, secretQuestion, secretAnswer) values (?,?,?,?)"
	_, err := db.Exec(sqlStr, u.username, u.password, u.secretQuestion, u.secretAnswer)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
}
func updateRowDemo(db *sql.DB,password string,username string) {
	sqlStr := "update user set password=? where username = ?"
	_, err := db.Exec(sqlStr, password, username)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
}
func check(r *gin.Engine,db *sql.DB){
	r.POST("/find", func(c *gin.Context) {
		username:=c.PostForm("username")
		answer:=c.PostForm("answer")
		u,flag:=QueryRowDemoUsername(db,username)
		if !flag {
			c.String(403,"无此账号")
			return
		}else {
			if u.secretAnswer == answer {
				c.String(200,"你的密码是%s",u.password)
			}
		}
	})
}