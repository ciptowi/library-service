package response

import (
	"github.com/labstack/echo/v4"
)

func ErrorResponse(c echo.Context, code int, message string, err error) error {
	res := new(RespError)
	res.Message = message
	res.Error = err
	return c.JSON(code, res)
}

func ErrorLogin(c echo.Context, code int, message string, token string) error {
	res := new(RespLogin)
	res.Message = message
	res.AccessToken = token
	return c.JSON(code, res)
}
