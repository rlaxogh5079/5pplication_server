package server

import (
	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()

	router.GET("/users", GETAllUsers)
	router.GET("/article/:long/:lat/", GETArticles)
	router.GET("login", GETLogin)
	router.POST("/login", POSTLogin)

	return router
}
