package server

import (
	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/controllers"
	"github.com/quyenphamkhac/go-tinyurl/repos"
	"github.com/quyenphamkhac/go-tinyurl/services"
)

func Serve() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api")
	{
		urls := api.Group("/urls")
		{
			urlRepo := new(repos.URLRespository)
			urlService := services.NewUrlService(urlRepo)
			urlCtrl := controllers.NewURLController(urlService)
			urls.GET("/", urlCtrl.GetAllURLs)
		}
	}

	r.Run()
}
