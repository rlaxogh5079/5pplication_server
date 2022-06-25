package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// @Summary get Image in some imageNo
// @ID load-image
// @Produce png
// @Tags image
// @name load-image
// @Success 200 {object} string "message" : "some-message"
// @Router /image [get]
// @Param imageNo query string true "imageNo"
func GETImage(c *gin.Context) {
	imageNo := c.Request.URL.Query()["imageNo"]
	if imageNo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "err : unauthorized",
		})
		fmt.Println("잘못된 접근입니다.")
		return
	}
	img, openErr := os.Open(fmt.Sprintf("/images/%v.png", imageNo[0]))
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

// @Summary insert Image in some imageNo
// @ID upload-image
// @Produce json
// @Tags image
// @name upload-image
// @Success 200 {object} string "message" : "some-message"
// @Router /image/upload [post]
// @Param imageNo formData file true "imageNo"
func POSTImageUpload(c *gin.Context) {
	// Single file
	image, loadErr := c.FormFile("imageNo")
	if loadErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err : imageNotExist",
		})
		fmt.Println("이미지가 존재하지 않습니다.")
		return
	}
	if _, err := os.Stat(fmt.Sprintf("/images/%v", image.Filename)); err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "err : imageAlreadyExist",
		})
		fmt.Println("이미지가 이미 존재합니다.")
		return
	}
	c.SaveUploadedFile(image, filepath.Join("/images", image.Filename))

	c.JSON(http.StatusOK, gin.H{
		"message": "imageSuccessfullyUpload",
	})
	fmt.Println("이미지가 성공적으로 업로드 되었습니다.")
}
