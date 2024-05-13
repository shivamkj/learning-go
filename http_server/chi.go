package http_server

import (
	"encoding/json"
	"fmt"
	net_http "net/http"

	"github.com/go-chi/chi/v5"
)

type chiContext struct {
	*net_http.Request
}

func (c chiContext) GetQueryParam(key string) string {
	return c.URL.Query().Get(key)
}

func (c chiContext) GetPathParam(key string) string {
	return chi.URLParam(c.Request, key)
}

func (c chiContext) ParseBody(i any) {
	defer c.Request.Body.Close()
	_ = json.NewDecoder(c.Request.Body).Decode(i)
}

func _(endpoints []HttpEndpoint, port int16) {
	muxObj := chi.NewRouter()

	for _, element := range endpoints {
		switch handler := element.Handler; element.Method {
		case GET:
			muxObj.Get(element.Path, func(w net_http.ResponseWriter, r *net_http.Request) {
				response, _ := json.Marshal(map[string]any{
					"success": true,
					"body":    handler(chiContext{r}),
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(net_http.StatusOK)
				w.Write(response)
			})
		case POST:
			muxObj.Post(element.Path, func(w net_http.ResponseWriter, r *net_http.Request) {
				response, _ := json.Marshal(map[string]any{
					"success": true,
					"body":    handler(chiContext{r}),
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(net_http.StatusOK)
				w.Write(response)
			})
		case PUT:
		case DELETE:
			fmt.Println("Currently not supported")
		}
	}

	err := net_http.ListenAndServe(fmt.Sprintf(":%d", port), muxObj)
	if err != nil {
		panic(err)
	}
}
