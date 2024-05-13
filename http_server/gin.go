package http_server

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ginContext struct {
	*gin.Context
}

func (c ginContext) GetQueryParam(key string) string {
	param, exists := c.GetQuery(key)
	if !exists {
		return ""
	}
	return param
}

func (c ginContext) GetPathParam(key string) string {
	return c.Param(key)
}

func (c ginContext) ParseBody(i any) {
	defer c.Request.Body.Close()
	_ = json.NewDecoder(c.Request.Body).Decode(i)
}

func _(endpoints []HttpEndpoint, port int16) {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)

	for _, element := range endpoints {
		switch handler := element.Handler; element.Method {
		case GET:
			engine.GET(element.Path, func(c *gin.Context) {
				result := handler(ginContext{c})
				c.JSON(200, gin.H{
					"message": result,
				})
			})
		case POST:
			engine.POST(element.Path, func(c *gin.Context) {
				result := handler(ginContext{c})
				c.JSON(200, gin.H{
					"message": result,
				})
			})
		case PUT:
		case DELETE:
			fmt.Println("Currently not supported")
		}
	}

	err := engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
