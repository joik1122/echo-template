package route

import (
	apiTest "echo-template/route/test"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

//Init 메인에 echo객체를 반환
func Init(e *echo.Echo) {

	e.GET("/healthz", func(c echo.Context) error {
		/**
		 * 필수입니다. 쿠버네티스 probe가 healthz url에 200을 성공적으로 받지 못하면 POD이 죽은걸로 판단합니다.
		 * 관련설정 /k8s/templates/deployment.yaml에서 확인
		 */
		return c.String(http.StatusOK, "Health checked....")
	})

	apiTest.Route(e.Group("/test"))

}
