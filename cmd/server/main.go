package main

import (
	"fluxmach-speed-engine/internal/config"
	"fluxmach-speed-engine/internal/handlers"
	"fluxmach-speed-engine/internal/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	// Load configuration
	cfg := config.DefaultConfig()

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/info", handlers.InfoHandler)
	mux.HandleFunc("/ping", handlers.PingHandler)
	mux.HandleFunc("/download", handlers.DownloadHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// Wrap mux with middleware
	handler := middleware.SetupMiddleware(mux, cfg.AuthKey)

	// Server setup
	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Fluxmach Speed Engine starting on %s...", cfg.Port)
	log.Printf("Routes: /info, /ping, /download, /upload, /health")
	
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", cfg.Port, err)
	}
}
