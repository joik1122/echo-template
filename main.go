package main

import (
	"echo-template/lib/err"
	"echo-template/lib/util"
	"echo-template/route"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Request에 대한 표준출력 콘솔 로그
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: util.UrlSkipper,
	}))
	// 모든 위치에서 panic을 복구하고 스택 추적을 인쇄하고 HTTPErrorHandler에 대한 제어를 처리
	e.Use(middleware.Recover())
	// GZIP을 통해 HTTP 응답을 압축
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 9, // 압축 레벨
	}))

	// CORS 설정
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	// 에러 핸들러 설정
	e.HTTPErrorHandler = err.ApiHTTPErrorHandler
	// validator 설정
	e.Validator = &CustomValidator{validator: validator.New()}
	// 라우팅 설정
	route.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
