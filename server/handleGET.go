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
		c.JSON(http.StatusInternalServerError, gin.H {
			"message" : longErr,
		})
	}
	lat, latErr := strconv.ParseFloat(c.Param("lat"), 64)
	if latErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message" : latErr,
		})
	}
	articleData, loadErr := database.LoadArticle(long, lat)
	if loadErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message" : loadErr,
		})
	}
	c.JSON(http.StatusOK, articleData)
}

func GETAllUsers(c *gin.Context) {
	users, loadErr := database.LoadUsers()
	if loadErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message" : loadErr,
		})
	}
	c.JSON(http.StatusOK, users)
}
