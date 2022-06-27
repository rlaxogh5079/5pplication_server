package server

import (
	"fmt"
	"net/http"

	database "github.com/5pplication/Server/db"
	login "github.com/5pplication/Server/login"
	"github.com/gin-gonic/gin"
)

// @Summary login in same email and hashed_password
// @ID login
// @Produce json
// @Tags user
// @name login
// @Success 200 {object} string "message" : "some-message"
// @Router /login [post]
// @param Email header string true "Email"
// @param Password header string true "Password"
func POSTLogin(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	fmt.Printf("%v님이 접속을 시도합니다.\n", email)
	password := c.Request.Header["Password"][0]
	db, connectErr := database.ConnectDB()

	if connectErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", connectErr.Error()),
		})
		fmt.Println(connectErr.Error())
		return
	}
	defer db.Close()

	result, loginErr := login.CheckLogin(db, email, password)
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

// @Summary signup in some database.User
// @ID signup
// @Produce json
// @Tags user
// @name signup
// @Success 200 {object} string "message" : "some-message"
// @Router /signup [post]
// @Param user body database.User true "user"
func POSTSignUp(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	nickname := c.Request.Header["Nickname"][0]
	password := c.Request.Header["Password"][0]
	storeArticle := "{}"
	db, connectErr := database.ConnectDB()

	if connectErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", connectErr.Error()),
		})
		fmt.Println(connectErr.Error())
		return
	}
	defer db.Close()

	flag, insertErr := database.InsertUser(db, email, nickname, password, storeArticle)
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

// @Summary delete user in some email and hashed_password
// @ID delete_user
// @Produce json
// @Tags user
// @name delete_user
// @Success 200 {object} string "message" : "some-message"
// @Router /user/delete [post]
// @param Email header string true "Email"
// @param Password header string true "Password"
func POSTDeleteUser(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	password := c.Request.Header["Password"][0]
	db, connectErr := database.ConnectDB()

	if connectErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", connectErr.Error()),
		})
		fmt.Println(connectErr.Error())
		return
	}
	defer db.Close()
	result, loginErr := login.CheckLogin(db, email, password)

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
	} else if loginErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", loginErr.Error()),
		})
		fmt.Println(loginErr.Error())
		return
	} else {
		flag, removeErr := database.RemoveUser(db, email)
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

// @Summary update user password in some email and hashed_password
// @ID update_user_password
// @Produce json
// @Tags user
// @name update_user_password
// @Success 200 {object} string "message" : "some-message"
// @Router /user/update/password [post]
// @param Email header string true "Email"
// @param Password header string true "Password"
func POSTUpdatePassword(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	password := c.Request.Header["Password"][0]
	db, connectErr := database.ConnectDB()

	if connectErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", connectErr.Error()),
		})
		fmt.Println(connectErr.Error())
		return
	}
	defer db.Close()

	result, updateErr := database.UpdatePassword(db, email, password)
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

// @Summary update user password in some email and nickname
// @ID update_user_nickname
// @Produce json
// @Tags user
// @name update_user_nickname
// @Success 200 {object} string "message" : "some-message"
// @Router /user/update/nickname [post]
// @param Email header string true "Email"
// @param Nickname header string true "Nickname"
func POSTUpdateNickname(c *gin.Context) {
	email := c.Request.Header["Email"][0]
	nickname := c.Request.Header["Nickname"][0]
	db, connectErr := database.ConnectDB()

	if connectErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", connectErr.Error()),
		})
		fmt.Println(connectErr.Error())
		return
	}
	defer db.Close()

	result, updateErr := database.UpdateNickname(db, email, nickname)
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
