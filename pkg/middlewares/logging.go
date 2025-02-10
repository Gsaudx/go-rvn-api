package middlewares

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the incoming request details.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s completed in %v", r.Method, r.URL.Path, time.Since(start))
	})
	//TODO Implement the logging of the request details
}
