package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	database "github.com/5pplication/Server/db"
	"github.com/gin-gonic/gin"
)

func GETArticles(c *gin.Context) {
	atclNo := c.Request.URL.Query()["atclNo"]
	long := c.Request.URL.Query()["long"]
	lat := c.Request.URL.Query()["lat"]

	if atclNo != nil {
		article, loadErr := database.SelectArticle(atclNo[0])
		if loadErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("err : %v", loadErr.Error()),
			})
			fmt.Println(loadErr.Error())
			return
		} else {
			c.JSON(http.StatusOK, article)
		}
	} else if long != nil && lat != nil {
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
		articles, loadErr := database.LoadArticle(long2, lat2)
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
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "err : unauthorized",
		})
	}
}

func POSTInsertArticle(c *gin.Context) {
	atclNo := c.Request.Header["Atclno"][0]
	email := c.Request.Header["Email"][0]
	share, parseBoolErr := strconv.ParseBool(c.Request.Header["Share"][0])
	long := c.Request.Header["Long"][0]
	lat := c.Request.Header["Lat"][0]
	title := c.Request.Header["Title"][0]
	body := c.Request.Header["Body"][0]
	date := string(time.Now().UTC().Format("2006-01-02 15:04:05"))
	images := c.Request.Header["Images"][0]
	tag := c.Request.Header["Tag"][0]

	flag, insertErr := database.InsertArticle(atclNo, email, share, long, lat, title, body, 0, date, images, tag)

	if parseBoolErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", parseBoolErr.Error()),
		})
		fmt.Println(parseBoolErr.Error())
		return
	}
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", insertErr.Error()),
		})
		fmt.Println(insertErr.Error())
		return
	}
	if flag {
		flag2, updateErr := database.UpdateStoreArticle(email, atclNo)
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

func POSTDeleteArticle(c *gin.Context) {
	atclNo := c.Request.Header["Atclno"][0]
	email := c.Request.Header["Email"][0]

	article, loadErr := database.SelectArticle(atclNo)
	fmt.Println(email, article.Email)
	if loadErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("err : %v", loadErr.Error()),
		})
		fmt.Println(loadErr.Error())
		return
	}

	if article.Email == email {
		flag, removeErr := database.RemoveArticle(atclNo)
		if removeErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("err : %v", removeErr.Error()),
			})
			fmt.Println(removeErr.Error())
			return
		}
		if flag {
			user, selectErr2 := database.SelectUser(email)
			if selectErr2 != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("err : %v", selectErr2.Error()),
				})
				fmt.Println(selectErr2.Error())
				return
			}
			flag2, removeErr2 := database.RemoveStoreArticle(user.Email, atclNo)
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
