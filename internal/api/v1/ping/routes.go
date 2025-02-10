package ping

import (
	"net/http"

	"github.com/gsaudx/api-go-project/pkg/middlewares"
)

// RegisterRoutes registers ping endpoint routes.
func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/v1/ping", middlewares.SecurityMiddleware(http.HandlerFunc(PingHandler)))
}
