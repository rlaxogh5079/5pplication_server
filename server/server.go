package server

import (
	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()

	router.GET("/article/:long/:lat/", GETArticles)
	router.POST("/login", POSTLogin)
	router.POST("/signup", POSTSignUp)
	router.POST("/delete/user", POSTDeleteUser)
	router.POST("/user/update/password", POSTUpdatePassword)
	router.POST("/user/update/nickname", POSTUpdateNickname)

	return router
}
