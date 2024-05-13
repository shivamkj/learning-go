package http_server

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type fiberContext struct {
	*fiber.Ctx
}

func (c fiberContext) GetQueryParam(key string) string {
	return c.Query(key)
}

func (c fiberContext) GetPathParam(key string) string {
	return c.Params(key)
}

func (c fiberContext) ParseBody(i any) {
	json.Unmarshal(c.BodyRaw(), i)
}

func handleRequest(handler HttpHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		response := handler(fiberContext{c})
		c.Send(response)
		c.Response().Header.SetContentType(fiber.MIMEApplicationJSON)
		return nil
	}
}

func handleError(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Send custom error page
	err = ctx.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   "Custom Error",
	})

	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).SendString("My Internal Server Error")
	}
	return nil
}

func Boot(endpoints []HttpEndpoint, port int16) {
	fiberInstance := fiber.New(fiber.Config{
		ErrorHandler: handleError, // Override default error handler
	})

	for _, element := range endpoints {
		switch handler := element.Handler; element.Method {
		case GET:
			fiberInstance.Get(element.Path, handleRequest(handler))
		case HEAD:
			fiberInstance.Head(element.Path, handleRequest(handler))
		case POST:
			// fiberInstance.Post(element.Path, handleRequest(handler))
		case PUT:
			fiberInstance.Post(element.Path, handleRequest(handler))
		case DELETE:
			fiberInstance.Post(element.Path, handleRequest(handler))
		case OPTIONS:
			fiberInstance.Post(element.Path, handleRequest(handler))
		case PATCH:
			fiberInstance.Post(element.Path, handleRequest(handler))

		}
	}

	err := fiberInstance.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
