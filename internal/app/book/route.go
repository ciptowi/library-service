package book

import (
	"github.com/labstack/echo/v4"
	"library-sevice/internal/middlewares"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create, middlewares.IsAuthenticated)
	g.PUT("/:id", h.UpdateByID, middlewares.IsAuthenticated)
	g.DELETE("/:id", h.DeleteByID, middlewares.IsAuthenticated)
}
