package server

import (
	"net/http"
	"strconv"

	database "github.com/5pplication/Server/db"
	"github.com/gin-gonic/gin"
)

func GETArticles(c *gin.Context) {
	long, longErr := strconv.ParseFloat(c.Param("long"), 64)
	if longErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": longErr.Error(),
		})
	}
	lat, latErr := strconv.ParseFloat(c.Param("lat"), 64)
	if latErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": latErr.Error(),
		})
	}
	articleData, loadErr := database.LoadArticle(long, lat)
	if loadErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": loadErr.Error(),
		})
	}
	c.JSON(http.StatusOK, articleData)
}
