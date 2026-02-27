package handlers

import (
	"encoding/json"
	"fluxmach-speed-engine/internal/config"
	"net"
	"net/http"
	"strings"
)

type InfoResponse struct {
	IP         string `json:"ip"`
	ISP        string `json:"isp"`
	ServerName string `json:"server_name"`
	Location   string `json:"location"`
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.DefaultConfig()

	// Try to get real IP from headers if behind a proxy
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			ip = host
		} else {
			ip = r.RemoteAddr
		}
	}

	// Handle multiple IPs in X-Forwarded-For
	if strings.Contains(ip, ",") {
		ip = strings.TrimSpace(strings.Split(ip, ",")[0])
	}

	// For local testing, return a more "realistic" ISP
	isp := "Local Provider"
	if ip == "127.0.0.1" || ip == "::1" {
		isp = "Local Network"
	}

	resp := InfoResponse{
		IP:         ip,
		ISP:        isp,
		ServerName: cfg.ServerName,
		Location:   cfg.Location,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
