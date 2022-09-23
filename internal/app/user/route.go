package user

import (
	"github.com/labstack/echo/v4"
	"library-sevice/internal/middlewares"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get, middlewares.IsAuthenticated)
	g.GET("/:id", h.GetByID, middlewares.IsAuthenticated)
	g.POST("", h.Create)
	g.PUT("/:id", h.UpdateByID, middlewares.IsAuthenticated)
	g.DELETE("/:id", h.DeleteByID, middlewares.IsAuthenticated)
}
