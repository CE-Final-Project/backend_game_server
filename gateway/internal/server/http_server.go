package server

import (
	"github.com/ce-final-project/backend_game_server/gateway/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
	"time"

	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	maxHeaderBytes = 1 << 20
	stackSize      = 1 << 10 // 1 KB
	bodyLimit      = "1M"
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	gzipLevel      = 5
)

func (s *Server) runHttpServer() error {
	s.mapRoutes()

	s.echo.Server.ReadTimeout = readTimeout
	s.echo.Server.WriteTimeout = writeTimeout
	s.echo.Server.MaxHeaderBytes = maxHeaderBytes

	return s.echo.Start(s.cfg.HTTP.Port)
}

func (s *Server) mapRoutes() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "API Gateway Game Server"
	docs.SwaggerInfo.Description = "API Gateway microservices."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	s.echo.Use(s.mw.RequestLoggerMiddleware)
	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))
	s.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	s.echo.Use(middleware.BodyLimit(bodyLimit))
}
