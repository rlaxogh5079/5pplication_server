package server

import (
	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()

	router.GET("/users", GETAllUsers)
	router.GET("/article/:long/:lat/", GETArticles)
	router.POST("/login", POSTLogin)
	router.POST("/signup", POSTSignUp)
	router.POST("/del/user", POSTDelUser)

	return router
}
