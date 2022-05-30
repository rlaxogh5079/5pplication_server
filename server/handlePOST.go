package server

import (
	"fmt"
	"net/http"

	database "github.com/5pplication/Server/db"
	login "github.com/5pplication/Server/login"
	"github.com/gin-gonic/gin"
)

func POSTLogin(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	fmt.Printf("%v님이 접속을 시도합니다.\n", email)
	password := []byte(c.Request.Header["Password"][0])
	result, loginErr := login.CheckLogin(email, password)
	if result == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   loginErr.Error(),
			"message": "아이디가 존재하지 않습니다.",
		})
		fmt.Println("아이디가 존재하지 않습니다.")
	} else if result == 0 { // 비밀번호가 틀림
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "비밀번호가 틀렸습니다.",
		})
		fmt.Println("비밀번호가 틀렸습니다.")
	} else { // 로그인 성공
		c.JSON(http.StatusOK, gin.H{
			"message": "환영합니다.",
		})
		fmt.Println("환영합니다.")
	}
}

func POSTSignUp(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	nickname := c.Request.Header["Nickname"][0]
	password := (c.Request.Header["Password"][0])
	storeArticle := "{}"
	insertErr := database.InsertUser(email, nickname, password, storeArticle)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   insertErr.Error(),
			"message": "해당 이메일은 이미 가입되었습니다.",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "회원가입 성공",
		})
		fmt.Printf("%v님이 회원가입 하였습니다.\n", email)
	}
}

func POSTDelUser(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	password := []byte(c.Request.Header["Password"][0])
	result, loginErr := login.CheckLogin(email, password)
	if loginErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   loginErr.Error(),
			"message": "로그인할 수 없습니다",
		})
	}
	if result == -1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "아이디가 존재하지 않습니다.",
		})
	} else if result == 0 { // 비밀번호가 틀렸을경우
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "비밀번호가 틀렸습니다.",
		})
	} else {
		database.RemoveUser(email)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v 계정이 성공적으로 제거되었습니다.", email),
		})
	}
}
