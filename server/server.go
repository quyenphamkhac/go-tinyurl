package server

import (
	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/config"
)

func Serve() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}

	r.Run(config.GetConfig().ServerConfig.Port)
}
