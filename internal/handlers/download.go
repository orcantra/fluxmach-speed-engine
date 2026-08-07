package handlers

import (
	"crypto/rand"
	"net/http"
	"time"
)

// DownloadHandler streams random bytes for a fixed duration
// Default duration: 10 seconds
// Endpoint: GET /download?duration=10s
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse duration
	durationStr := r.URL.Query().Get("duration")
	duration := 10 * time.Second // Default
	if durationStr != "" {
		if d, err := time.ParseDuration(durationStr); err == nil {
			duration = d
		}
	}

	// Set headers for streaming
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Transfer-Encoding", "chunked")
	// Cache-Control is set by middleware

	// Flush the headers to start the response immediately
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	// Stream data
	// Small buffer so chunks arrive frequently even on slow links - a large
	// buffer can take many seconds to fill, making the transfer look stalled.
	buffer := make([]byte, 32*1024)

	// Pre-fill buffer with random data
	rand.Read(buffer)

	flusher, canFlush := w.(http.Flusher)

	timeout := time.After(duration)

	for {
		select {
		case <-timeout:
			return
		case <-r.Context().Done():
			// Client disconnected
			return
		default:
			_, err := w.Write(buffer)
			if err != nil {
				return
			}
			// Flush every chunk so a proxy in front can't buffer past `duration`.
			if canFlush {
				flusher.Flush()
			}
		}
	}
}
