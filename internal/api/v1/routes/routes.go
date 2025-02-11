package routes

import (
	"net/http"

	"github.com/gsaudx/go-tasks-api/internal/api/v1/ping"
	// Import additional endpoints here when needed, e.g.:
	// "github.com/gsaudx/go-tasks-api/internal/api/v1/users"
)

// RegisterRoutes aggregates and registers all v1 API endpoint routes.
func RegisterRoutes(mux *http.ServeMux) {
	ping.RegisterRoutes(mux)
	// Add additional route registrations here, e.g.:
	// users.RegisterRoutes(mux)
}
