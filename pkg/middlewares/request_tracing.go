package middlewares

import (
	"net/http"

	"github.com/google/uuid"
)

const RequestIDHeader = "X-Request-ID"

// RequestTracingMiddleware adds a unique request ID to each request.
func RequestTracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		w.Header().Set(RequestIDHeader, requestID)
		// Optionally, you can store the requestID in the context:
		// ctx := context.WithValue(r.Context(), "requestID", requestID)
		// r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
