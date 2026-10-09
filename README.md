# Go E-Commerce Backend API

A hands-on backend learning project built with Go. The application currently provides an in-memory product API and will gradually evolve into a database-backed e-commerce service.

## Current status

**Lessons 1–4 completed; Lesson 5 in progress:** HTTP server, structs and slices, GET/POST product endpoints, JSON validation, and fetching a product by ID. Concurrency-safety improvements are being reviewed.

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

The latest shared `main.go` listens on **http://localhost:8080**. If port 8080 is occupied, change `ListenAndServe` to `:8081` and adjust the examples below.

## API reference

| Method | Endpoint | Description | Status |
| --- | --- | --- | --- |
| `GET` | `/products` | List all products | Implemented |
| `POST` | `/products` | Create a product | Implemented |
| `GET` | `/products/{id}` | Get one product by numeric ID | Implemented |
| `PUT` | `/products/{id}` | Update a product | In progress (not yet implemented) |
| `DELETE` | `/products/{id}` | Delete a product | Planned |
| `POST` | `/auth/register` | Register user | Planned |
| `POST` | `/auth/login` | Log in | Planned |
| `POST` | `/orders` | Create order | Planned |

### GET /products

```bash
curl -i http://localhost:8080/products
```

Example response (`200 OK`):

```json
[
  {"id":1,"name":"Laptop","price":899.99,"stock":20},
  {"id":2,"name":"Mouse","price":19.99,"stock":50},
  {"id":3,"name":"Keyboard","price":5.99,"stock":200}
]
```

### POST /products

```bash
curl -i -X POST http://localhost:8080/products \
  -H 'Content-Type: application/json' \
  -d '{"name":"Gaming Mouse","price":49.99,"stock":15}'
```

Example response (`201 Created`, assuming the three initial products):

```json
{"id":4,"name":"Gaming Mouse","price":49.99,"stock":15}
```

The server assigns the product ID; clients do not need to supply one.

### Input validation

The create endpoint rejects:

- Invalid JSON, unknown fields, or multiple JSON values
- Empty or whitespace-only product names
- Negative prices (zero is currently accepted)
- Negative stock quantities

Invalid requests return `400 Bad Request`. The last shared code uses `1<<30` (1 GiB) despite a 1 MB comment; change this to `1<<20` for a 1 MiB limit. This is basic validation for learning, not comprehensive production validation.

Test an invalid request:

```bash
curl -i -X POST http://localhost:8080/products \
  -H 'Content-Type: application/json' \
  -d '{"name":"","price":-5,"stock":-1}'
```

### GET /products/{id}

Fetch one product by its numeric ID:

```bash
curl -i http://localhost:8080/products/2
```

Example response (`200 OK`):

```json
{"id":2,"name":"Mouse","price":19.99,"stock":50}
```

Invalid IDs such as `/products/abc` or `/products/-1` return `400 Bad Request`. A valid ID that does not exist, such as `/products/999`, returns `404 Not Found`.

The handler reads the route parameter with `r.PathValue("id")`, converts it using `strconv.Atoi`, and searches the in-memory slice.

### Lesson 5: PUT /products/{id} (in progress)

Planned request example (not yet a confirmed working endpoint):

```bash
curl -i -X PUT http://localhost:8080/products/2 \
  -H 'Content-Type: application/json' \
  -d '{"name":"Gaming Mouse","price":29.99,"stock":30}'
```

The next task is to implement the update handler, validate its JSON body, safely modify the matching product, and return `404` when the product is absent.

## Concepts learned

- Go modules, packages, imports, and `main()`
- Structs, exported fields, and JSON tags
- Slices, `append`, and maps
- Functions and HTTP handlers
- `http.ServeMux` routing
- `json.NewEncoder` and `json.NewDecoder`
- Pointers (`&input`) and error handling
- Validation and HTTP status codes (`200`, `201`, `400`)
- `sync.Mutex` for protecting shared state (full handler coverage still needs verification)
- Path parameters with `r.PathValue("id")`
- String-to-integer conversion with `strconv.Atoi`
- Searching slices with `for ... range`, `break`, and `404 Not Found`

## Roadmap

- [x] **Lesson 1:** Initialize a Go module and start an HTTP server
- [x] **Lesson 2:** Define product structs and implement `GET /products`
- [x] **Lesson 3:** Implement `POST /products`, JSON decoding, validation, and basic mutex protection
- [x] **Lesson 4:** Implement `GET /products/{id}` and handle `400` / `404` (quiz: 4/4)
- [ ] **Lesson 5:** Implement `PUT /products/{id}` (in progress; not yet confirmed working)
- [ ] **Lesson 6:** Implement `DELETE /products/{id}`
- [ ] **Lesson 8:** Organize packages, handlers, services, and repositories
- [ ] **Lesson 9:** PostgreSQL integration and migrations
- [ ] **Lesson 10:** Authentication and authorization
- [ ] **Lesson 11:** Shopping cart and order transactions
- [ ] **Lesson 12:** Goroutines, channels, and background jobs
- [ ] **Lesson 13:** Context, timeouts, and graceful shutdown
- [ ] **Lesson 14:** Redis caching
- [ ] **Lesson 15:** Unit and integration testing
- [ ] **Lesson 16:** Docker and Compose
- [ ] **Lesson 17:** CI/CD and deployment

## Development commands

```bash
go fmt ./...                 # Format code
go vet ./...                 # Static checks
go test ./...                # Run tests (as they are added)
go build -o bin/api ./cmd/api # Build executable
```

## Code review follow-ups

- Change `var found product` to `var found Product` in the lesson 4 handler (Go is case-sensitive).
- Guard **all** reads and writes of the shared `products` slice with the same mutex, including `GET /products` and `POST /products`.
- Avoid holding the mutex while writing HTTP responses; copy the needed product under the lock and unlock before encoding.
- Change the POST request body limit from `1<<30` to `1<<20` if the intended limit is 1 MiB.
- Use `"GET /"` (with a space) for the home route; `"GET/"` is not the intended method-qualified pattern.
- Verify the final code with `go build ./...` and `go test -race ./...` when tests are added.

## Known limitations

- Products are stored in memory and disappear after restart.
- The last shared POST handler generates IDs with `len(products)+1`; this should be replaced with a synchronized counter or database-generated IDs.
- `float64` is used for prices for now; production monetary values should use integer minor units or a decimal type.
- There is no authentication, persistence, pagination, or automated test suite yet.
- Mutex usage is not yet consistent across all handlers in the latest shared code; it does not provide persistence or multi-instance coordination.

## License

No license has been selected yet.
