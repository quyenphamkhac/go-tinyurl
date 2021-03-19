package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/controllers"
	"github.com/quyenphamkhac/go-tinyurl/db"
	"github.com/quyenphamkhac/go-tinyurl/repos"
	"github.com/quyenphamkhac/go-tinyurl/services"
)

var (
	session *gocql.Session = db.GetDb()
)

func Serve() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api")
	{
		urls := api.Group("/urls")
		{
			urlRepo := repos.NewURLRepository(session)
			urlService := services.NewUrlService(urlRepo)
			urlCtrl := controllers.NewURLController(urlService)
			urls.GET("/", urlCtrl.GetAllURLs)
		}
	}

	r.Run()
}
