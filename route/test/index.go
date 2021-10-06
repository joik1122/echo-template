package test

import "github.com/labstack/echo/v4"

func Route(g *echo.Group) {
	g.GET("", getTest)
}
