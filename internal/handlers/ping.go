package handlers

import (
	"net/http"
)

// PingHandler returns an empty response for latency measurement
// No compression, no cache, very small payload
func PingHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
