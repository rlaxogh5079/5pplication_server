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

func POSTInsertArticle(c *gin.Context) {
	atclNo := c.Request.Header["Atclno"][0]
	email := c.Request.Header["Email"][0]
	share, parseBoolErr := strconv.ParseBool(c.Request.Header["Share"][0])
	long := c.Request.Header["Long"][0]
	lat := c.Request.Header["Lat"][0]
	title := c.Request.Header["Title"][0]
	body := c.Request.Header["Body"][0]
	date := string(time.Now().UTC().Format("2006-01-02 15:04:05"))
	tag := c.Request.Header["Tag"][0]

	insertErr := database.InsertArticle(atclNo, email, share, long, lat, title, body, 0, date, tag)

	if parseBoolErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": parseBoolErr.Error(),
		})
		fmt.Println(parseBoolErr.Error())
		return
	}

	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": insertErr.Error(),
		})
		fmt.Println(insertErr.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successInsertArticle",
	})
	fmt.Println("성공적으로 글을 입력하였습니다.")
}

func POSTDeleteArticle(c *gin.Context) {
	atclNo := c.Request.Header["Atclno"][0]
	email := c.Request.Header["Email"][0]

	article, loadErr := database.SelectArticle(atclNo)
	fmt.Println(email, article.Email)
	if loadErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": loadErr.Error(),
		})
		fmt.Println(loadErr.Error())
		return
	}

	if article.Email == email {
		flag, removeErr := database.RemoveArticle(atclNo)
		if removeErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": removeErr.Error(),
			})
			fmt.Println(removeErr.Error())
			return
		}

		if flag {
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
}
