package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Hello World!",
		"length":  23,
	})
}
