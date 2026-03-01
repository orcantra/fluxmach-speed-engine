package middleware

import (
	"net/http"
)

// SetupMiddleware wraps the handler with necessary middleware
func SetupMiddleware(next http.Handler, authKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Cache-Control, X-Fluxmach-Key")

		// No Cache
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Health check is public
		if r.URL.Path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		// API Key Validation
		key := r.Header.Get("X-Fluxmach-Key")
		if key != authKey {
			http.Error(w, "Unauthorized: Invalid or missing API Key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
