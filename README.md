# Go Tasks API 🚀

A lightweight and efficient task management API built with Go, featuring proper middleware chaining, Docker support, and RESTful design principles.

![Go Version](https://img.shields.io/badge/go-1.23.6-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

## Features ✨

- JWT Authentication (scaffolded - see `internal/middlewares/auth.go`)
- Request Tracing with UUIDs
- Structured Logging
- Security Headers (CSP, XSS Protection)
- Graceful Shutdown
- Docker & Docker Compose Support
- Standardized JSON Responses

## Getting Started 🏁

### Prerequisites
- Go 1.23+
- Docker (optional)

### Installation
```bash
git clone https://github.com/yourusername/go-tasks-api.git
cd go-tasks-api
go mod download
```

### Running locally with Go
```bash
go run cmd/api/main.go
```

### Running with Docker
```bash
docker-compose up --build
```

## API Endpoints 📡

### Health Check
```bash
curl https://localhost:8080/v1/ping
```

Example Reponse:
```json
{
"status": "success",
"code": 200,
"message": "Request successful",
"data": {
    "message": "pong"
    }
}
```

## Project Structure 📂
```
.
├── cmd/
│ └── api/ # Main application entrypoint
├── internal/
│ ├── api/ # API route handlers
│ ├── config/ # Environment configuration
│ └── middlewares/ # Application-specific middleware
├── pkg/
│ └── middlewares/ # Reusable middleware components
└── Dockerfile # Multi-stage Docker build
```

## Configuration 🔧
Environment variables:
```bash
PORT=8080 # Server Port (default: 8080)
```

## Middleware Chain 🔗
The application uses the following middleware sequence:
1. Request Tracing (UUID)
2. Security Headers
3. Authentication
4. Logging

## Contributing 🤝
Contributions are welcome! Please ensure:
1. Add tests for new features
2. Maintain consistent middleware patterns
3. Use the standardized response format