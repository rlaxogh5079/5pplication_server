package rest

import (
	"net/http"
	"strconv"

	database "github.com/5pplication/Server/db"
	"github.com/gin-gonic/gin"
)

func HandleArticles(c *gin.Context) {
	long, longerr := strconv.ParseFloat(c.Param("long"), 6)
	if longerr != nil {
		panic(longerr)
	}
	lat, laterr := strconv.ParseFloat(c.Param("lat"), 6)
	if laterr != nil {
		panic(laterr)
	}
	articleData, loadErr := database.LoadArticle(long, lat)
	if loadErr != nil {
		panic(loadErr)
	}
	c.JSON(http.StatusOK, articleData)
}
