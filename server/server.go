package server

import (
	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()

	router.POST("/login", POSTLogin)
	router.POST("/signup", POSTSignUp)
	router.POST("/user/delete", POSTDeleteUser)
	router.POST("/user/update/password", POSTUpdatePassword)
	router.POST("/user/update/nickname", POSTUpdateNickname)

	router.GET("/article", GETArticles)
	router.POST("/article/insert", POSTInsertArticle)
	router.POST("/article/delete", POSTDeleteArticle)

	return router
}
