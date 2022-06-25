package server

import (
	docs "github.com/5pplication/Server/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Api() *gin.Engine {

	router := gin.Default()
	docs.SwaggerInfo.Title = "5pplication API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is a 5pplication Server Swagger"
	docs.SwaggerInfo.BasePath = "/"

	router.POST("/login", POSTLogin)
	router.POST("/signup", POSTSignUp)
	router.POST("/user/delete", POSTDeleteUser)
	router.POST("/user/update/password", POSTUpdatePassword)
	router.POST("/user/update/nickname", POSTUpdateNickname)

	router.GET("/article/get", GETArticles)
	router.POST("/article/insert", POSTInsertArticle)
	router.POST("/article/delete", POSTDeleteArticle)

	router.GET("/image", GETImage)
	router.POST("/image/upload", POSTImageUpload)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
