package router

import (
	"github.com/gin-gonic/gin"
	"go-redis/api"
)

func InitRouter() {
	router := gin.Default()
	router.POST("/api/code", api.CodeHandle)
	router.POST("/api/login", api.LoginHandle)


	router.Run(":9090")
}