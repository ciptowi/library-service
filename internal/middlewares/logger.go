package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
	"os"
)

var CostumLogger = middleware.LoggerConfig{
	Format:           " ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri} ",
	CustomTimeFormat: "2006/01/02 15:04:05",
	Output:           os.Stdout,
}
