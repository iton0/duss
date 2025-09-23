## Project Structure
```
duss/
├── api-gateway-service
│   ├── cmd
│   │   └── server
│   │       └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   └── internal
│       ├── api
│       │   ├── handlers.go
│       │   └── handlers_test.go
│       ├── core
│       │   └── services
│       │       ├── api_gateway.go
│       │       └── api_gateway_test.go
│       └── infrastructure
│           ├── clients
│           │   ├── shortener_client.go
│           │   └── redirect_client.go
│           └── web
│               ├── router.go
│               └── router_test.go
├── ARCHITECTURE.md
├── CODE_OF_CONDUCT.md
├── CONTRIBUTING.md
├── docker-compose.yml
├── go.work
├── go.work.sum
├── key-gen-service
│   ├── cmd
│   │   └── server
│   │       └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   └── internal
│       ├── api
│       │   ├── handlers.go
│       │   └── handlers_test.go
│       ├── core
│       │   └── services
│       │       ├── key_generator.go
│       │       └── key_generator_test.go
│       └── infrastructure
│           └── web
│               ├── router.go
│               └── router_test.go
├── LICENSE
├── Makefile
├── mise.toml
├── README.md
├── scripts
│   └── init.sh
├── shared
│   ├── domain
│   │   └── url.go
│   └── go.mod
├── timeline.txt
├── url-redirect-service
│   ├── cmd
│   │   └── server
│   │       └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── internal
│       ├── api
│       │   ├── handlers.go
│       │   └── handlers_test.go
│       ├── core
│       │   └── services
│       │       ├── redirect.go
│       │       └── redirect_test.go
│       └── infrastructure
│           ├── storage
│           │   ├── mock
│           │   │   └── mock_storage.go
│           │   ├── postgres.go
│           │   ├── postgres_test.go
│           │   ├── redis.go
│           │   ├── redis_test.go
│           │   └── storage.go
│           └── web
│               ├── router.go
│               └── router_test.go
└── url-shortener-service
    ├── cmd
    │   └── server
    │       └── main.go
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    └── internal
        ├── api
        │   ├── handlers.go
        │   └── handlers_test.go
        ├── core
        │   └── services
        │       ├── shortener.go
        │       └── shortener_test.go
        └── infrastructure
            ├── storage
            │   ├── mock
            │   │   └── mock_storage.go
            │   ├── postgres.go
            │   ├── postgres_test.go
            │   ├── redis.go
            │   ├── redis_test.go
            │   └── storage.go
            └── web
                ├── router.go
                └── router_test.go
```
---

### Root Directory (`duss`)

This directory is a monorepo, holding all the microservices and shared project-level files.

- **docker-compose.yml:** The primary orchestration file. It is now responsible for defining and running multiple services: the three Go applications (url-shortener, url-redirect, key-gen), a PostgreSQL container, and a Redis container.

---

### API Gateway Service (`api-gateway-service`)

This service is the **only one exposed to the public internet**. All client requests must go through it.

- **cmd/server/main.go:** The entry point. It's responsible for orchestrating the entire gateway by initializing and connecting its internal components.
- **internal/api/handlers.go:** Contains the public-facing HTTP handlers. It does not contain business logic; instead, it delegates requests to the core gateway service.
- **internal/core/services/api_gateway.go:** The core business logic for the gateway. It implements the `GatewayServiceIface` and contains the orchestration logic to delegate requests to the correct internal client.
- **internal/infrastructure/clients:** It contains the concrete HTTP client implementations that know how to communicate with the other services on the internal network.
- **internal/infrastructure/web/router.go:** Defines the public API endpoints that the outside world will use.

---

### Backend Services (`url-shortener-service`, `url-redirect-service`, `key-gen-service`)

These services are **internal**. They are only accessible within the private Docker network and should not be exposed on public ports.

- **cmd/server/main.go:** The entry point for this specific microservice. It initializes and runs only the components required for its single purpose.
- **internal/api/handlers.go:** The handlers for this service's internal API endpoints. They are called by the API gateway's clients, not by external clients.
- **internal/core/services:** The core business logic for this specific service. It is completely isolated.
- **internal/infrastructure/storage:** Contains the storage implementations (PostgreSQL, Redis) needed by this service.
- **internal/infrastructure/web:** This package is responsible for all internal web-facing concerns. The router.go file defines and initializes the Gin router for its internal API endpoints.

---

### Shared Directory

This directory contains data models and interfaces that are common to multiple microservices.

- **shared/domain/url.go:** Defines the `URL` data structure used by multiple services.

---

### API Endpoints

The API is now unified under the `api-gateway-service`.

#### 1. Shorten a URL

The client sends a request to the gateway.

- **Public-facing URL:** `api-gateway-service.com/shorten`
- **Method:** `POST`
- **Functionality:** The `api-gateway-service` receives the request and **internally** calls the `url-shortener-service`'s API to perform the shortening. The short URL is then returned to the client.

#### 2. Redirect to Original URL

The client sends a request for the short key to the gateway.

- **Public-facing URL:** `api-gateway-service.com/:shortKey`
- **Method:** `GET`
- **Functionality:** The `api-gateway-service` receives the request and **internally** calls the `url-redirect-service`'s API to get the original URL. The gateway then issues a 301 redirect response to the client.

#### 3. Generate a Short Key

This endpoint is **only** used internally by the `url-shortener-service`.

- **Internal URL:** `key-gen-service.com/api/v1/generate-key`
- **Method:** `GET`
- **Functionality:** The `url-shortener-service` calls this endpoint to get a unique key for a new URL.
