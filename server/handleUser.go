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
	password := c.Request.Header["Password"][0]
	result, loginErr := login.CheckLogin(email, password)
	if loginErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", loginErr.Error()),
		})
		fmt.Println(loginErr.Error())
		return
	}
	if result == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "noId",
		})
		fmt.Println("Id가 존재하지 않습니다.")
	} else if result == 0 { // 비밀번호가 틀림
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrongPassword",
		})
		fmt.Println("비밀번호가 틀렸습니다.")
	} else { // 로그인 성공
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome",
		})
		fmt.Println("환영합니다.")
	}
}

func POSTSignUp(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	nickname := c.Request.Header["Nickname"][0]
	password := c.Request.Header["Password"][0]
	storeArticle := "{}"
	flag, insertErr := database.InsertUser(email, nickname, password, storeArticle)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", insertErr.Error()),
		})
		fmt.Println(insertErr.Error())
		return
	}
	if flag {
		c.JSON(http.StatusOK, gin.H{
			"message": "successInsert",
		})
		fmt.Println("회원가입을 성공했습니다.")
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "failedInsert",
		})
		fmt.Println("회원가입을 실패했습니다.")
	}
}

func POSTDeleteUser(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	password := c.Request.Header["Password"][0]
	result, loginErr := login.CheckLogin(email, password)
	if loginErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", loginErr.Error()),
		})
		fmt.Println(loginErr.Error())
		return
	}
	if result == -1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "noId",
		})
		fmt.Println("Id가 존재하지 않습니다.")
	} else if result == 0 { // 비밀번호가 틀렸을경우
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrongPassword",
		})
		fmt.Println("비밀번호가 틀렸습니다.")
	} else {
		flag, removeErr := database.RemoveUser(email)
		if removeErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("err : %v", removeErr.Error()),
			})
			fmt.Println(removeErr.Error())
			return
		}
		if flag {
			c.JSON(http.StatusOK, gin.H{
				"message": "fsuccessDelete",
			})
			fmt.Printf("%v가 성공적으로 제거되었습니다.\n", email)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "failedDelete",
			})
			fmt.Printf("%v가 제거되지 않았습니다.\n", email)
		}
	}
}

func POSTUpdatePassword(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	password := c.Request.Header["Password"][0]
	result, updateErr := database.UpdatePassword(email, password)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", updateErr.Error()),
		})
		return
	}
	if result {
		c.JSON(http.StatusOK, gin.H{
			"message": "successUpdate",
		})
		fmt.Println("성공적으로 업데이트 되었습니다.")
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "cantUpdate",
		})
		fmt.Println("업데이트 할 수 없습니다.")
	}
}

func POSTUpdateNickname(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	nickname := c.Request.Header["Nickname"][0]
	result, updateErr := database.UpdateNickname(email, nickname)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", updateErr.Error()),
		})
		fmt.Println(updateErr.Error())
		return
	}
	if result {
		c.JSON(http.StatusOK, gin.H{
			"message": "successUpdate",
		})
		fmt.Println("성공적으로 업데이트 되었습니다.")
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "cantUpdate",
		})
		fmt.Println("업데이트 할 수 없습니다.")
	}
}
