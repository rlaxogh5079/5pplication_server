package rest

import (
	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()

	router.GET("/users", HandleAllUsers)
	router.GET("/article/:long/:lat/", HandleArticles)

	return router
}