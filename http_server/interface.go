package http_server

type (
	HttpHandler func(Request) []byte

	HttpMethod int

	HttpEndpoint struct {
		Method  HttpMethod
		Path    string
		Handler HttpHandler
	}

	Request interface {
		GetQueryParam(key string) string
		GetPathParam(key string) string
		ParseBody(i any)
	}
)

const (
	GET HttpMethod = iota
	HEAD
	POST
	PUT
	DELETE
	OPTIONS
	PATCH
	// CONNECT
	// TRACE
)
