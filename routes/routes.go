package routes

import (
	"manifesto/config"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"

	manifestHandler "manifesto/domains/manifest/handler"
	manifestRepo "manifesto/domains/manifest/repositories"
	manifestService "manifesto/domains/manifest/service"
)

func InitRoutes(e *echo.Echo, rediscli *redis.Client, cfg *config.AppConfig) {
	/*
		Dependency Injection
	*/
	manifestrepo := manifestRepo.New(rediscli)
	manifestservice := manifestService.New(manifestrepo)
	manifesthandler := manifestHandler.New(manifestservice)

	/*
		Routes
	*/

	e.POST("/manifest", manifesthandler.Post)
	e.GET("/manifest/:key", manifesthandler.Get)
	e.GET("/:key", manifesthandler.Redirect)
}
