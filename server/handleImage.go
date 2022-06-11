package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GETImage(c *gin.Context) {
	image := c.Request.URL.Query()["imageNo"]
	if image == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "err : unauthorized",
		})
		fmt.Println("잘못된 접근입니다.")
		return
	}
	img, openErr := os.Open(fmt.Sprintf("images/%v.png", image[0]))
	if openErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err : imageNotExist",
		})
		fmt.Println("이미지가 존재하지 않습니다.")
		return
	}
	imgData, _ := ioutil.ReadAll(img)

	c.Data(http.StatusOK, "image/png", imgData)
}
