package user

import (
	"github.com/labstack/echo/v4"
	"library-sevice/internal/dto"
	"library-sevice/internal/factory"
	"library-sevice/pkg/utils/response"
	"net/http"
	"strconv"
)

type handler struct {
	service *service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Create(c echo.Context) error {
	payload := new(dto.CreateUserRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Failed", err)
	}
	msg, err := h.service.Create(c.Request().Context(), *payload)
	if msg == "Email already exists!" {
		return response.ErrorResponse(c, http.StatusBadRequest, msg, err)
	}
	return response.SuccessOk(c, http.StatusCreated, msg, payload)
}

func (h *handler) Get(c echo.Context) error {
	payload := new(dto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Failed", err)
	}

	data, pages, err := h.service.Find(c.Request().Context(), payload)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "Failed", err)
	}

	return response.SuccessOkWithPagination(c, http.StatusOK, "Succes", data, pages)
}

func (h *handler) GetByID(c echo.Context) error {
	param := new(dto.ByIDRequest)
	if err := c.Bind(param); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Failed params name", err)
	}

	data, err := h.service.FindByID(c.Request().Context(), param)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "Failed", err)
	}
	if data.ID == 0 {
		var empty interface{}
		return response.SuccessOk(c, http.StatusNotFound, "Data empty", empty)
	}
	return response.SuccessOk(c, http.StatusOK, "Succes", data)
}

func (h *handler) UpdateByID(c echo.Context) error {
	//param := new(dto.ByIDRequest)
	id, _ := strconv.Atoi(c.Param("id"))
	//if err := c.Bind(param); err != nil {
	//	return response.ErrorResponse(c, http.StatusBadRequest, "Failed params name", err)
	//}

	payload := new(dto.UpdateUserRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Failed", err)
	}

	data, err := h.service.UpdateByID(c.Request().Context(), uint(id), *payload)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "Failed", err)
	}

	return response.SuccessOk(c, http.StatusOK, "Succes", data)
}
func (h *handler) DeleteByID(c echo.Context) error {
	param := new(dto.ByIDRequest)
	if err := c.Bind(param); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "Failed params name", err)
	}

	data, err := h.service.DeleteByID(c.Request().Context(), param)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "Failed", err)
	}

	return response.SuccessOk(c, http.StatusOK, "Succes", data)
}
