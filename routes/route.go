package routes

import (
	"github.com/fachrunwira/basic-template-go-echo/controllers/user"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", user.Home)
}
