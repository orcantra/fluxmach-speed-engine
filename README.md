# Fluxmach Speed Engine

A lightweight, high-performance network speed measurement engine written in Go.

Maintained and published by **[Orcantra](https://orcantra.com)**.

## Prerequisites

- [Go](https://go.dev/dl/) 1.21 or higher.

## Setup

1.  Clone the repository (or ensure files are in place).
2.  Run the server:
    ```bash
    go run cmd/server/main.go
    ```

## Usage

The server listens on port `8080`.

### 1. Connection Info

**Endpoint:** `GET /info`

- Returns JSON containing client IP, ISP information, and server metadata (name/location).
- Requires `X-Fluxmach-Key` header.
- Used to display connection context on the dashboard.
- Example:
  ```json
  {
    "ip": "127.0.0.1",
    "isp": "Local Network",
    "server_name": "Fluxmach Edge Node 1",
    "location": "New York, US"
  }
  ```

### 2. Latency (Ping)

**Endpoint:** `GET /ping`

- Returns 200 OK with empty body.
- Requires `X-Fluxmach-Key` header.
- Used to measure RTT and jitter.

### 3. Download Speed

**Endpoint:** `GET /download`

- Optional parameters: `duration` (e.g., `10s`, default: `10s`).
- Requires `X-Fluxmach-Key` header.
- Streams random data for the specified duration.
- Example: `curl -H "X-Fluxmach-Key: your-key" http://localhost:8080/download?duration=5s > /dev/null`

### 4. Upload Speed

**Endpoint:** `POST /upload`

- Accepts any data stream and discards it.
- Requires `X-Fluxmach-Key` header.
- Example: `curl -X POST -H "X-Fluxmach-Key: your-key" --data-binary @largefile http://localhost:8080/upload`

### 5. Health Check

**Endpoint:** `GET /health`

- Returns 200 OK.

## License

This project is [MIT licensed](LICENSE).
