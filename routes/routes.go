package routes

import (
	"immersive/config"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *redis.Client, cfg *config.AppConfig) {
	/*
		Dependency Injection
	*/

	/*
		Routes
	*/

}
