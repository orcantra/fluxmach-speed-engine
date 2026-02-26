package handlers

import (
	"io"
	"net/http"
)

// UploadHandler accepts data stream and discards it
// Endpoint: POST /upload
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Discard the request body
	_, err := io.Copy(io.Discard, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
}
