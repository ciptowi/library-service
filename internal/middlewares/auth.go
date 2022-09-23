package middlewares

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
)

var JWT_KEY = os.Getenv("JWT_SECRET_KEY")

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(JWT_KEY),
})
