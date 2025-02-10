package middlewares

import "net/http"

// AuthMiddleware checks for a valid authorization token.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// token := r.Header.Get("Authorization")
		// if token != "Bearer Token" { // Replace with your own validation logic
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }
		next.ServeHTTP(w, r)
	})
}
