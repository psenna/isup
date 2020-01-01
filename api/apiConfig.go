package api

import (
	"github.com/gin-gonic/gin"
	"github.com/psenna/isup/api/controllers"
)

func GetApiServer() *gin.Engine {
	server := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	server.GET("/", controllers.GetHome())

	return server
}
