package routes

import (
	"net/http"

	"github.com/gsaudx/api-go-project/internal/api/v1/ping"
	//TODO Import additional endpoints here when needed
)

// RegisterRoutes aggregates and registers all v1 API endpoint routes.
func RegisterRoutes(mux *http.ServeMux) {
	ping.RegisterRoutes(mux)
	//TODO Add additional route registrations here, e.g.:
	//TODO users.RegisterRoutes(mux)
}
