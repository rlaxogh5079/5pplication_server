package server

import (
	"fmt"
	"net/http"

	login "github.com/5pplication/Server/login"
	"github.com/gin-gonic/gin"
)

func POSTLogin(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	email := c.Request.Header["Email"][0]
	fmt.Println(email)
	password := []byte(c.Request.Header["Password"][0])
	result, _ := login.CheckLogin(email, password)
	c.String(http.StatusOK, fmt.Sprintf("%v", result))
}

func GETLogin(c *gin.Context) {

}
