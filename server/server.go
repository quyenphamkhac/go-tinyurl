package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/config"
	"github.com/quyenphamkhac/go-tinyurl/controllers"
	"github.com/quyenphamkhac/go-tinyurl/datasources"
	"github.com/quyenphamkhac/go-tinyurl/middlewares"
	"github.com/quyenphamkhac/go-tinyurl/repos"
	"github.com/quyenphamkhac/go-tinyurl/services"
)

func Serve() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	session := datasources.GetSession()
	jwtService := services.NewJwtService(time.Hour*2, config.GetConfig().Secret, config.GetConfig().Issuer)

	api := r.Group("/api")
	{
		urls := api.Group("/urls")
		urls.Use(middlewares.AuthorizeWithJwt(jwtService))
		{
			urlRepo := repos.NewURLRepository(session)
			urlService := services.NewUrlService(urlRepo)
			urlCtrl := controllers.NewURLController(urlService)
			urls.GET("/", urlCtrl.GetAllURLs)
			urls.POST("/", urlCtrl.CreateURL)
		}

		auth := api.Group("/auth")
		{
			userRepo := repos.NewUserRepository(session)
			authService := services.NewAuthService(userRepo, jwtService)
			authCtrl := controllers.NewAuthController(authService)
			auth.POST("/signup", authCtrl.SignUp)
			auth.POST("/login", authCtrl.Login)
		}
	}

	r.Run()
}
