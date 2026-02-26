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

### 1. Latency (Ping)

**Endpoint:** `GET /ping`

- Returns 200 OK with empty body.
- Used to measure RTT and jitter.

### 2. Download Speed

**Endpoint:** `GET /download`

- Optional parameters: `duration` (e.g., `10s`, default: `10s`).
- Streams random data for the specified duration.
- Example: `curl http://localhost:8080/download?duration=5s > /dev/null`

### 3. Upload Speed

**Endpoint:** `POST /upload`

- Accepts any data stream and discards it.
- Example: `curl -X POST --data-binary @largefile http://localhost:8080/upload`

### 4. Health Check

**Endpoint:** `GET /health`

- Returns 200 OK.

## License

This project is [MIT licensed](LICENSE).
