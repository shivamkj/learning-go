package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type echoContext struct {
	echo.Context
}

func (c echoContext) GetQueryParam(key string) string {
	return c.QueryParam(key)
}

func (c echoContext) GetPathParam(key string) string {
	return c.Param(key)
}

func (c echoContext) ParseBody(i any) {
	defer c.Request().Body.Close()
	_ = json.NewDecoder(c.Request().Body).Decode(i)
}

func _(endpoints []HttpEndpoint, port int) {
	echoInstance := echo.New()

	for _, element := range endpoints {
		switch handler := element.Handler; element.Method {
		case GET:
			echoInstance.GET(element.Path, func(c echo.Context) error {
				return c.JSON(http.StatusOK, handler(echoContext{c}))
			})
		case POST:
			echoInstance.POST(element.Path, func(c echo.Context) error {
				return c.JSON(http.StatusOK, handler(echoContext{c}))
			})
		case PUT:
		case DELETE:
			fmt.Println("Currently not supported")
		}
	}

	err := echoInstance.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
