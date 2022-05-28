package server

import (
	"fmt"
	"net/http"

	login "github.com/5pplication/Server/login"
	database "github.com/5pplication/Server/db"
	"github.com/gin-gonic/gin"
)

func POSTLogin(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	email := c.Request.Header["Email"][0]
	fmt.Println(email)
	password := []byte(c.Request.Header["Password"][0])
	result, loginErr := login.CheckLogin(email, password)
	if loginErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message" : loginErr,
		})
	}
	fmt.Println(result)
	if result == 0 { // 비밀번호가 틀림
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "비밀번호가 틀렸습니다.",
	})
	} else { // 로그인 성공
		c.JSON(http.StatusOK, gin.H {
			"message" : "환영합니다.",
		})
	}
}

func POSTSignUp(c* gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	email := c.Request.Header["Email"][0]
	nickname := c.Request.Header["Nickname"][0]
	password := (c.Request.Header["Password"][0])
	storeArticle := "{}"
	insertErr := database.InsertUser(email, nickname, password, storeArticle)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message" : insertErr,
		})
	} else {
		c.String(http.StatusOK, "회원가입 성공")
	}
}