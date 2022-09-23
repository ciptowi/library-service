package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"library-sevice/internal/app/auth"
	"library-sevice/internal/app/book"
	"library-sevice/internal/app/user"
	"library-sevice/internal/factory"
	"library-sevice/internal/middlewares"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Use(middleware.LoggerWithConfig(middlewares.CostumLogger))

	v1 := e.Group("/api/v1")
	auth.NewHandler(f).Route(v1.Group(""))
	book.NewHandler(f).Route(v1.Group("/books"))
	user.NewHandler(f).Route(v1.Group("/users"))
}
