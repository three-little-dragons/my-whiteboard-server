package main

import (
	"github.com/gin-gonic/gin"
	"github.com/three-little-dragons/my-whiteboard-server/internal/app/http"
)

func setupRouter() *gin.Engine {
	//gin.SetMode(config.Val.Gin.LogLevel)
	r := gin.Default()
	initRouter(r)
	return r
}

func initRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/snowflake_id", http.SnowflakeId)
	}
}
