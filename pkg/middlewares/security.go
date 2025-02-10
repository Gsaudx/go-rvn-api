package middlewares

import "net/http"

// SecurityMiddleware adds standard security headers to the response.
func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevent MIME sniffing.
		w.Header().Set("X-Content-Type-Options", "nosniff")
		// Enable basic XSS protection.
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		// Set Content Security Policy.
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		next.ServeHTTP(w, r)
	})
}
