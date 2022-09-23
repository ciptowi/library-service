package auth

import (
	"github.com/labstack/echo/v4"
	"library-sevice/internal/dto"
	"library-sevice/internal/factory"
	"library-sevice/pkg/utils/response"
	"net/http"
)

type handler struct {
	service *service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

func (h *handler) Login(c echo.Context) error {

	payload := new(dto.AuthLoginRequest)
	if err = c.Bind(payload); err != nil {
		return err
	}

	token, data, err := h.service.Login(c.Request().Context(), payload)
	if err != nil {
		return response.ErrorLogin(c, http.StatusForbidden, "Wrong password!", "")
	}

	return response.SuccessLogin(c, http.StatusOK, "Login Success", data, token)
}
