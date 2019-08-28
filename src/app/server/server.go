package server

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"app/controller"
	v9 "app/utils/validator"
	"app/utils/response"
	"app/config"
	"app/router"
)

// App is a data structure to keep the defined application.
type App struct {
	echo    *echo.Echo
	conf    *config.Config
	server  *http.Server
}

// Initialise the app with predefined configuration.
func Initialise(conf *config.Config) error {
	app := &App{
		conf: conf,
	}

	app.echo = echo.New()

	// Middlewares
	app.echo.Use(middleware.Logger())
	app.echo.Use(middleware.JWTWithConfig(router.JWT(conf)))
	app.echo.Use(middleware.CORSWithConfig(router.CORS()))
	app.echo.Use(middleware.Secure())

	app.server = &http.Server{
		Addr:         app.conf.Bind,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Setup core handlers
	c := controller.New(conf, v9.New(), response.New())

	// Routes
	app.echo.POST("/auth/register", c.Register)
	app.echo.POST("/auth/login", c.Login)

	// Initialise
	app.echo.Logger.Fatal(app.start())
	return nil
}

// Start fires up the HTTP server.
func (app *App) start() error {
	return app.echo.StartServer(app.server)
}