# Go E-Commerce Backend API

A hands-on backend learning project built with Go. The application currently provides an in-memory product API and will gradually evolve into a database-backed e-commerce service.

## Current status

**Lessons 1–3 completed:** HTTP server, structs and slices, GET/POST product endpoints, JSON decoding, basic validation, and mutex-protected in-memory storage.

> This is a learning project, not a production-ready store. Data is reset when the server restarts.

## Tech stack

| Technology | Purpose | Status |
| --- | --- | --- |
| Go 1.26.2 | Backend language | In use |
| `net/http` / `http.ServeMux` | HTTP server and routing | In use |
| `encoding/json` | JSON request/response handling | In use |
| `sync.Mutex` | Protect shared in-memory products | In use |
| PostgreSQL | Persistent database | Planned |
| JWT | Authentication | Planned |
| Redis | Cache / background processing | Planned |
| Docker | Containerization | Planned |
| GitHub Actions | CI/CD | Planned |

## Project structure

```text
go-ecommerce/
├── cmd/
│   └── api/
│       └── main.go
├── go.mod
└── README.md
```

## Prerequisites

- Go 1.26 or newer (developed with Go 1.26.2 on Linux/amd64)
- A terminal and a code editor

## Run locally

From the project root:

```bash
go fmt ./...
go build ./...
go run ./cmd/api
```

The server currently listens on **http://localhost:8081**. Port 8081 is used because port 8080 was occupied on the development machine.

## API reference

| Method | Endpoint | Description | Status |
| --- | --- | --- | --- |
| `GET` | `/products` | List all products | Implemented |
| `POST` | `/products` | Create a product | Implemented |
| `GET` | `/products/{id}` | Get one product | Next lesson |
| `PUT` | `/products/{id}` | Update a product | Planned |
| `DELETE` | `/products/{id}` | Delete a product | Planned |
| `POST` | `/auth/register` | Register user | Planned |
| `POST` | `/auth/login` | Log in | Planned |
| `POST` | `/orders` | Create order | Planned |

### GET /products

```bash
curl -i http://localhost:8081/products
```

Example response (`200 OK`):

```json
[
  {"id":1,"name":"Laptop","price":999.99,"stock":10},
  {"id":2,"name":"Keyboard","price":79.99,"stock":25}
]
```

### POST /products

```bash
curl -i -X POST http://localhost:8081/products \
  -H 'Content-Type: application/json' \
  -d '{"name":"Gaming Mouse","price":49.99,"stock":15}'
```

Example response (`201 Created`, assuming the two initial products):

```json
{"id":3,"name":"Gaming Mouse","price":49.99,"stock":15}
```

The server assigns the product ID; clients do not need to supply one.

### Input validation

The create endpoint rejects:

- Invalid JSON, unknown fields, or multiple JSON values
- Empty or whitespace-only product names
- Prices less than or equal to zero
- Negative stock quantities

Invalid requests return `400 Bad Request`. The request body is limited to approximately 1 MiB. This is basic validation for learning, not comprehensive production validation.

Test an invalid request:

```bash
curl -i -X POST http://localhost:8081/products \
  -H 'Content-Type: application/json' \
  -d '{"name":"","price":-5,"stock":-1}'
```

## Concepts learned

- Go modules, packages, imports, and `main()`
- Structs, exported fields, and JSON tags
- Slices, `append`, and maps
- Functions and HTTP handlers
- `http.ServeMux` routing
- `json.NewEncoder` and `json.NewDecoder`
- Pointers (`&input`) and error handling
- Validation and HTTP status codes (`200`, `201`, `400`)
- Basic concurrency safety with `sync.Mutex`

## Roadmap

- [x] **Lesson 1:** Initialize a Go module and start an HTTP server
- [x] **Lesson 2:** Define product structs and implement `GET /products`
- [x] **Lesson 3:** Implement `POST /products`, JSON decoding, validation, and basic mutex protection
- [ ] **Lesson 4:** Implement `GET /products/{id}` and handle `404 Not Found`
- [ ] **Lesson 5:** Update and delete products
- [ ] **Lesson 6:** Organize packages, handlers, services, and repositories
- [ ] **Lesson 7:** PostgreSQL integration and migrations
- [ ] **Lesson 8:** Authentication and authorization
- [ ] **Lesson 9:** Shopping cart and order transactions
- [ ] **Lesson 10:** Goroutines, channels, and background jobs
- [ ] **Lesson 11:** Context, timeouts, and graceful shutdown
- [ ] **Lesson 12:** Redis caching
- [ ] **Lesson 13:** Unit and integration testing
- [ ] **Lesson 14:** Docker and Compose
- [ ] **Lesson 15:** CI/CD and deployment

## Development commands

```bash
go fmt ./...                 # Format code
go vet ./...                 # Static checks
go test ./...                # Run tests (as they are added)
go build -o bin/api ./cmd/api # Build executable
```

## Known limitations

- Products are stored in memory and disappear after restart.
- IDs are assigned by a local counter, not a database sequence.
- `float64` is used for prices for now; production monetary values should use integer minor units or a decimal type.
- There is no authentication, persistence, pagination, or automated test suite yet.
- A mutex protects shared product state within one process; it does not provide persistence or multi-instance coordination.

## License

No license has been selected yet.
