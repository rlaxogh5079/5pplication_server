package server

import (
	"net/http"
	"strconv"

	database "github.com/5pplication/Server/db"
	"github.com/gin-gonic/gin"
)

func GETArticles(c *gin.Context) {
	long, longerr := strconv.ParseFloat(c.Param("long"), 64)
	if longerr != nil {
		panic(longerr)
	}
	lat, laterr := strconv.ParseFloat(c.Param("lat"), 64)
	if laterr != nil {
		panic(laterr)
	}
	articleData, loadErr := database.LoadArticle(long, lat)
	if loadErr != nil {
		panic(loadErr)
	}
	c.JSON(http.StatusOK, articleData)
}
