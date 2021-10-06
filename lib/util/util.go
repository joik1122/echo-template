package util

import (
	"strings"

	echo "github.com/labstack/echo/v4"
)

// UrlSkipper 헬스체크 url에 대한 스키퍼 (로그/트레이싱 스킵함)
func UrlSkipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/healthz") {
		return true
	}
	return false
}
