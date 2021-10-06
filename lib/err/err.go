package err

import (
	"log"
	"net/http"
	"runtime"
	"time"

	echo "github.com/labstack/echo/v4"
)

func ApiHTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	} else {
		msg = http.StatusText(code)
	}

	responseData := map[string]interface{}{
		"error_msg": msg,
	}

	c.JSON(http.StatusOK, responseData)
}

func CallError(code int, errMsg string, c echo.Context) {
	req := c.Request()
	_, fn, line, _ := runtime.Caller(1)

	log.Printf(`[error log] time: "%s", code : "%d", method : "%s", uri : "%s", file: "%s", line: "%d", message: "%s"`, time.Now().Format(time.RFC3339), code, req.Method, req.RequestURI, fn, line, errMsg)

	panic(echo.NewHTTPError(code, errMsg))
}
