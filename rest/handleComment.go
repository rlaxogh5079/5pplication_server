package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/5pplication/Server/db"
)

func HandleComments(c* gin.Context){
	userId := c.Param("userId")
	comments := db.LoadComment(userId)
	c.JSON(http.StatusOK, comments)
}