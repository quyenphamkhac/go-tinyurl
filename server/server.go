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
	r.Use(middlewares.ErrorsMiddleware(gin.ErrorTypeAny))

	session := datasources.GetSession()
	jwtService := services.NewJwtService(time.Hour*2, config.GetConfig().Secret, config.GetConfig().Issuer)
	cache := datasources.GetRedisCache()
	cacheRepo := repos.NewCacheRepository(cache)

	urlRepo := repos.NewURLRepository(session, cacheRepo)
	urlService := services.NewUrlService(urlRepo)
	urlCtrl := controllers.NewURLController(urlService)

	userRepo := repos.NewUserRepository(session)
	authService := services.NewAuthService(userRepo, jwtService)
	authCtrl := controllers.NewAuthController(authService)

	main := r.Group("/tinyurl")
	{
		main.GET("/:hash", urlCtrl.RedirectURLByHash)
	}

	api := r.Group("/api")
	{
		urls := api.Group("/urls")
		urls.Use(middlewares.AuthorizeWithJwt(jwtService))
		{
			urls.GET("/:hash", urlCtrl.GetURLByHash)
			urls.POST("/", urlCtrl.CreateURL)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/signup", authCtrl.SignUp)
			auth.POST("/login", authCtrl.Login)
		}

	}

	r.Run()
}
