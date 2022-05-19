package rest

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/5pplication/Server/db"
)

func HandleArticles(c* gin.Context){
	long, longerr := strconv.ParseFloat(c.Param("long"),6)
	if longerr != nil{
		panic(longerr)
	}
	lat, laterr := strconv.ParseFloat(c.Param("lat"),6)
	if laterr != nil{
		panic(laterr)
	}
	articleData, loadErr := db.LoadArticle(long, lat)
	if loadErr != nil {
		panic(loadErr)
	}
	c.JSON(http.StatusOK, articleData)
}