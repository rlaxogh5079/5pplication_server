package server

import (
	"fmt"
	"net/http"
	"strconv"

	database "github.com/5pplication/Server/db"
	"github.com/gin-gonic/gin"
)

// @Summary get all articles in same long and lat
// @ID get-articles
// @Produce json
// @Tags article
// @name get-articles
// @Success 200 {object} database.Article
// @Router /article/get [get]
// @param atclNo query string false "atclNo"
// @Param long query string false "long"
// @Param lat query string false "lat"
func GETArticles(c *gin.Context) {
	atclNo := c.Request.URL.Query()["atclNo"]
	long := c.Request.URL.Query()["long"]
	lat := c.Request.URL.Query()["lat"]
	db, connectErr := database.ConnectDB()

	if connectErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": connectErr.Error(),
		})
		fmt.Println(connectErr.Error())
		return
	}
	defer db.Close()

	if atclNo != nil {
		article, loadErr := database.SelectArticle(db, atclNo[0])
		if loadErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("err : %v", loadErr.Error()),
			})
			fmt.Println(loadErr.Error())
			return
		} else {
			c.JSON(http.StatusOK, article)
		}
	}

	if long != nil && lat != nil {

		long2, longConvErr := strconv.ParseFloat(long[0], 64)
		if longConvErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("err : %v", longConvErr.Error()),
			})
			fmt.Println(longConvErr.Error())
			return
		}

		lat2, latConvErr := strconv.ParseFloat(lat[0], 64)
		if latConvErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("err : %v", latConvErr.Error()),
			})
			fmt.Println(latConvErr.Error())
			return
		}

		articles, loadErr := database.LoadArticle(db, long2, lat2)
		if loadErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("err : %v", loadErr.Error()),
			})
			fmt.Println(loadErr.Error())
			return
		} else {
			c.JSON(http.StatusOK, articles)
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "err : badrequest",
		})
	}
}

// @Summary Insert article in some database.Article
// @ID insert-article
// @Produce json
// @Tags article
// @name insert-article
// @Success 200 {object} string "message" : "some-message"
// @Router /article/insert [post]
// @Param article body database.Article true "article"
func POSTInsertArticle(c *gin.Context) {
	var article database.Article
	readErr := c.BindJSON(&article)
	if readErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", readErr.Error()),
		})
		fmt.Println(readErr.Error())
		return
	}
	atclNo := article.AtclNo
	email := article.Email
	share := article.Share
	long := article.Long
	lat := article.Lat
	title := article.Title
	body := article.Body
	date := article.Date
	images := article.Images
	tag := article.Tag
	db, connectErr := database.ConnectDB()

	if connectErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", connectErr.Error()),
		})
		fmt.Println(connectErr.Error())
		return
	}
	defer db.Close()

	flag, insertErr := database.InsertArticle(db, atclNo, email, share, long, lat, title, body, 0, date, images, tag)

	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", insertErr.Error()),
		})
		fmt.Println(insertErr.Error())
		return
	}

	if flag {
		flag2, updateErr := database.UpdateStoreArticle(db, email, atclNo)
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("err : %v", updateErr.Error()),
			})
			fmt.Println(updateErr.Error())
		}

		if flag2 {
			c.JSON(http.StatusOK, gin.H{
				"message": "successInsertArticle",
			})
			fmt.Println("성공적으로 글을 입력하였습니다.")
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "failedInsertArticle",
			})
			fmt.Println("글을 입력하는데 실패하였습니다.")
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "failedInsertArticle",
		})
		fmt.Println("글을 입력하는데 실패하였습니다.")
	}
}

// @Summary delete article in some atclno and email
// @ID delete-article
// @Produce json
// @Tags article
// @name delete-article
// @Success 200 {object} string "message" : "some-message"
// @Router /article/delete [post]
// @Param Atclno header string true "Atclno"
// @Param Email header string true "Email"
func POSTDeleteArticle(c *gin.Context) {
	atclNo := c.Request.Header["Atclno"][0]
	email := c.Request.Header["Email"][0]
	db, connectErr := database.ConnectDB()

	if connectErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", connectErr.Error()),
		})
		fmt.Println(connectErr.Error())
		return
	}
	defer db.Close()

	article, loadErr := database.SelectArticle(db, atclNo)
	fmt.Println(email, article.Email)
	if loadErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", loadErr.Error()),
		})
		fmt.Println(loadErr.Error())
		return
	}

	if article.Email == email {
		flag, removeErr := database.RemoveArticle(db, atclNo)
		if removeErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("err : %v", removeErr.Error()),
			})
			fmt.Println(removeErr.Error())
			return
		}
		if flag {
			user, selectErr2 := database.SelectUser(db, email)
			if selectErr2 != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("err : %v", selectErr2.Error()),
				})
				fmt.Println(selectErr2.Error())
				return
			}
			flag2, removeErr2 := database.RemoveStoreArticle(db, user.Email, atclNo)
			if removeErr2 != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("err : %v", removeErr2.Error()),
				})
				fmt.Println(removeErr2.Error())
				return
			}
			if flag2 {
				c.JSON(http.StatusOK, gin.H{
					"message": "successDelete",
				})
				fmt.Println("성공적으로 글을 제거하였습니다.")
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "failedDelete",
				})
				fmt.Println("글을 제거하지 못했습니다.")
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "emailDoesntMatch",
			})
			fmt.Println("이메일이 일치하지 않습니다.")
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "emailDoesntMatch",
		})
		fmt.Println("이메일이 일치하지 않습니다.")
	}
}
