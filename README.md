# Go E-Commerce Backend API

A hands-on Go backend development project built step by step to learn Go fundamentals, RESTful APIs, databases, authentication, concurrency, testing, and deployment.

## Tech Stack

- **Language:** Go 1.26+
- **HTTP Server:** Go Standard Library (`net/http`)
- **Router:** `http.ServeMux` (initial implementation)
- **Database:** PostgreSQL (planned)
- **Authentication:** JWT (planned)
- **Cache:** Redis (planned)
- **Containerization:** Docker (planned)
- **CI/CD:** GitHub Actions (planned)

## Project Structure

```text
go-ecommerce/
├── cmd/
│   └── api/
│       └── main.go
├── go.mod
└── README.md
```

Additional packages and directories will be introduced as the project grows.

## Getting Started

### Prerequisites

- Go 1.26 or newer
- Git
- A terminal
- VS Code or another code editor

### Installation

Clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/go-ecommerce.git
cd go-ecommerce
```

Replace `YOUR_USERNAME` with your GitHub username.

If working locally without a GitHub repository, simply open the existing `go-ecommerce` directory.

### Run the Server

```bash
go run ./cmd/api
```

The development server runs at:

```text
http://localhost:8081
```

### Test the API

For Lesson 1:

```bash
curl http://localhost:8081/
```

Expected response:

```json
{
  "message": "Welcome to Go E-Commerce API"
}
```

For Lesson 2, once implemented:

```bash
curl http://localhost:8081/products
```

## API Endpoints

| Method | Endpoint | Description | Status |
|---|---|---|---|
| GET | `/` | Welcome message | Implemented |
| GET | `/products` | List all products | In progress |
| GET | `/products/{id}` | Get product by ID | Planned |
| POST | `/products` | Create product | Planned |
| PUT | `/products/{id}` | Update product | Planned |
| DELETE | `/products/{id}` | Delete product | Planned |
| POST | `/auth/register` | Register user | Planned |
| POST | `/auth/login` | User login | Planned |
| POST | `/orders` | Create order | Planned |
| GET | `/orders` | Order history | Planned |

## Learning Roadmap

- [x] Lesson 1: Go project setup and first HTTP server
- [ ] Lesson 2: Structs, slices, and GET `/products`
- [ ] Lesson 3: JSON decoding and POST `/products`
- [ ] Lesson 4: Methods, pointers, and error handling
- [ ] Lesson 5: Packages and project organization
- [ ] Lesson 6: Interfaces and dependency injection
- [ ] Lesson 7: PostgreSQL integration
- [ ] Lesson 8: Database migrations and SQL queries
- [ ] Lesson 9: Authentication and authorization
- [ ] Lesson 10: Shopping cart and order management
- [ ] Lesson 11: Transactions and inventory management
- [ ] Lesson 12: Goroutines, channels, and background workers
- [ ] Lesson 13: Context and graceful shutdown
- [ ] Lesson 14: Redis caching
- [ ] Lesson 15: Unit and integration testing
- [ ] Lesson 16: Docker and Docker Compose
- [ ] Lesson 17: GitHub Actions and CI/CD
- [ ] Lesson 18: Production deployment

## Architecture (Planned)

```text
HTTP Client
    |
    v
Router / Middleware
    |
    v
Handler
    |
    v
Service
    |
    v
Repository
    |
    v
PostgreSQL
```

The initial implementation uses Go's standard library. Additional layers will be introduced when needed.

## Development Commands

Run the application:

```bash
go run ./cmd/api
```

Format Go code:

```bash
go fmt ./...
```

Check for common issues:

```bash
go vet ./...
```

Run tests:

```bash
go test ./...
```

Build the application:

```bash
go build -o bin/api ./cmd/api
```

## Learning Goals

By completing this project, the developer should be able to:

1. Understand Go syntax and its type system.
2. Build REST APIs using Go.
3. Work with structs, interfaces, pointers, and packages.
4. Handle errors and HTTP requests correctly.
5. Design database-backed backend applications.
6. Implement authentication and authorization.
7. Understand goroutines, channels, and concurrency.
8. Write automated tests.
9. Containerize and deploy Go applications.
10. Apply practical backend architecture principles.

## Project Status

**Current stage:** Lesson 1 completed, Lesson 2 in progress.

This is an educational project under active development.

## License

No license has been selected yet.
