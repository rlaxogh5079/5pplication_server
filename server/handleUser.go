package server

import (
	"net/http"

	database "github.com/5pplication/Server/db"
	"github.com/gin-gonic/gin"
)

func GETAllUsers(c *gin.Context) {
	users, loadErr := database.LoadUsers()
	if loadErr != nil {
		panic(loadErr)
	}
	c.JSON(http.StatusOK, users)
}
