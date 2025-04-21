# 📌 XuXaDex Backend Project

## 🚀 Overview
A high-performance, scalable backend service built with Golang. This project follows best practices for clean architecture, efficient database management, and robust API development.

## 🏗 Features
- 🌐 RESTful API with JSON responses
- 🗄️ PostgreSQL database integration
- 📜 Structured logging with Zap
- 📦 Dependency management using Go Modules
- 🧪 Unit & integration testing
- 🚀 Docker support for easy deployment
- 📡 API documentation via Swagger

## 🛠️ Tech Stack
- **Language**: Golang
- **Framework**: Echo (Fast & lightweight)
- **Database**: PostgreSQL
- **Logging**: Zap
- **Testing**: Go Testing

## ⚡ Quick Start

### Prerequisites
Ensure you have the following installed:
- [Go](https://golang.org/doc/install) (>=1.23)
- [Docker](https://docs.docker.com/get-docker/) (optional for containerized deployment)

### Installation
1. Clone the repository:
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up environment variables:
   ```sh
   cp .env.example .env
   ```
4. Run the application:
   ```sh
   go run main.go
   ```

### Running with Docker
```sh
docker-compose up --build
```

## 📖 API Documentation
API docs are available via Swagger:
```sh
make run
```
Visit `http://localhost:<your_ip>/api/v1/swagger/index.html` in your browser.

## 🧪 Running Tests
Run unit tests with:
```sh
go test ./...
```

## 🏗 Project Structure
```
/go-backend
│── /api           # Handlers & routes
│── /cmd           # Entry points (main.go)
│── /config        # Configuration files
│── /db            # Database migrations & queries
│── /doc           # Swagger documentation
│── /internal      # Business logic, services
│── /logs          # Logs directory
│── /pkg           # Utility functions & shared packages
│── /test          # Test files
└── go.mod         # Dependencies
```

## 🚀 Deployment
For production deployment, build the binary:
```sh
go build -o app ./cmd/main.go
```
Run the binary:
```sh
./app
```
