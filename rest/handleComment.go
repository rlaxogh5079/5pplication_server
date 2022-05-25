package rest

import (
	"net/http"

	database "github.com/5pplication/Server/db"
	"github.com/gin-gonic/gin"
)

func HandleComments(c *gin.Context) {
	userId := c.Param("userId")
	comments := database.LoadComment(userId)
	c.JSON(http.StatusOK, comments)
}
