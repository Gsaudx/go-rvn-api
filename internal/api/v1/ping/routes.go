package ping

import (
	"net/http"

	"github.com/gsaudx/api-go-project/pkg/middlewares"
)

// RegisterRoutes registers ping endpoint routes.
func RegisterRoutes(mux *http.ServeMux) {
	// For versioning, you could actually wrap the handler with a versioned prefix.
	mux.Handle("/v1/ping", middlewares.SecurityMiddleware(http.HandlerFunc(PingHandler)))
}
