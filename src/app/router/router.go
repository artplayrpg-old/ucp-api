package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"app/config"
)

func CORS() middleware.CORSConfig {
	return middleware.CORSConfig {
		AllowHeaders: []string{"X-Requested-With"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"HEAD", "GET", "POST", "PUT", "OPTIONS"},
	}
}

func JWT(conf *config.Config) middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: []byte(conf.JWT.Secret),
		Skipper: func(c echo.Context) bool {
			return true
		},
	}
}