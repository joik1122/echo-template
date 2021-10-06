package test

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getTest(c echo.Context) error {
	return c.String(http.StatusOK, "GET TEST!")
}
