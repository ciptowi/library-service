package response

import (
	"github.com/labstack/echo/v4"
	"library-sevice/internal/dto"
)

func SuccessOk(c echo.Context, code int, message string, data interface{}) error {
	res := new(RespOk)
	res.Success = true
	res.Message = message
	res.Data = data
	return c.JSON(code, res)
}

func SuccessOkWithPagination(c echo.Context, code int, message string, data interface{}, pagination dto.PaginationInfo) error {
	res := new(RespOkWithPagination)
	res.Success = true
	res.Message = message
	res.Data = data
	res.Pagiration = pagination
	return c.JSON(code, res)
}

func SuccessLogin(c echo.Context, code int, message string, data interface{}, token string) error {
	res := new(RespLogin)
	res.Success = true
	res.Message = message
	res.AccessToken = token
	res.TokenType = "Bearer"
	res.Data = data
	return c.JSON(code, res)
}
