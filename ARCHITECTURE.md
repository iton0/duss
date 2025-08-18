## Project Structure
```
duss/
├── shared/
│   └── domain/
│       └── url.go
├── url-shortener-service/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── api/
│   │   │   └── handlers.go
│   │   ├── core/
│   │   │   ├── services/
│   │   │   │   └── shortener.go
│   │   └── infrastructure/
│   │       ├── storage/
│   │       │   ├── mock/
│   │       │   │   └── mock_storage.go
│   │       │   ├── storage.go
│   │       │   └── postgres.go
│   │       └── web/
│   │           └── router.go
│   ├── go.mod
│   └── Dockerfile
├── url-redirect-service/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── api/
│   │   │   └── handlers.go
│   │   ├── core/
│   │   │   └── services/
│   │   │       └── redirect.go
│   │   └── infrastructure/
│   │       ├── storage/
│   │       │   ├── mock/
│   │       │   │   └── mock_storage.go
│   │       │   ├── storage.go
│   │       │   └── redis.go
│   │       └── web/
│   │           └── router.go
│   ├── go.mod
│   └── Dockerfile
├── key-gen-service/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── api/
│   │   │   └── handlers.go
│   │   ├── core/
│   │   │   └── services/
│   │   │       └── key_generator.go
│   │   └── infrastructure/
│   │       └── web/
│   │           └── router.go
│   │
│   ├── go.mod
│   └── Dockerfile
├── configs/
│   └── config.yaml
├── docker-compose.yml
└── scripts/
```
---

## Root Directory (`duss`)

This directory now acts as a monorepo, holding all the microservices and shared project-level files.

- **docker-compose.yml:** The primary orchestration file. It is now responsible for defining and running multiple services: the three Go applications (url-shortener, url-redirect, key-gen), a PostgreSQL container, and a Redis container.

## Service Directories (e.g., `url-shortener-service`)

Each service directory is a self-contained, independent Go module. It is a separate application with its own dependencies and executable.

- **cmd/server/main.go:** The entry point for this specific microservice. It is responsible for initializing and running only the components required for its single purpose.

- **internal/api/handlers.go:** Contains the HTTP handlers for this service's specific API endpoints. For the `url-shortener-service`, this would be the `POST /api/v1/shorten` handler. For the `url-redirect-service`, it would be the `GET /:shortKey` handler. The `key-gen-service` might expose an internal API for the shortener service to use.

- **internal/core:** The core business logic for this specific service. It is completely isolated from other services.

  - **internal/core/services:** Contains the business logic functions. For the `url-shortener-service`, this would handle the URL shortening process, including validating the input URL, requesting a key from the `key-gen-service`, and storing the mapping. For the `url-redirect-service`, this would handle the redirection lookup from the database.

- **internal/infrastructure:** Contains the tools and concrete implementations specific to this service.

  - **internal/infrastructure/storage:** Contains the storage implementations needed by this service. The `url-shortener-service` will use the PostgreSQL implementation, and the `url-redirect-service` will use the Redis implementation.

- **go.mod:** A separate Go module file for this service, allowing it to manage its dependencies independently.

- **Dockerfile:** A separate Dockerfile for building this service's container image.

## Shared Directory

This directory contains data models and interfaces that are common to multiple microservices.

- **shared/domain/url.go:** Defines the `URL` data structure used by both the `url-shortener-service` and the `url-redirect-service`. By centralizing this definition, it ensures consistency and prevents code duplication across services.

## API Endpoints

The API is now split across different services.

### 1. Shorten a URL

Creates a new short URL from a long URL. This endpoint is handled by the **`url-shortener-service`**.

- **URL:** `/api/v1/shorten`
- **Method:** `POST`
- **Functionality:** Accepts a long URL, validates it, requests a unique key from the `key-gen-service`, and saves the mapping to the PostgreSQL database.

### 2. Redirect to Original URL

Redirects a user from the short URL to the original long URL. This endpoint is handled by the **`url-redirect-service`**.

- **URL:** `/:shortKey`
- **Method:** `GET`
- **Functionality:** Looks up the `shortKey` in the Redis cache (or PostgreSQL as a fallback) and issues a 301 redirect to the original long URL.

### 3. Generate a Short Key

Generates a new, unique short key for a URL. This internal endpoint is handled by the **`key-gen-service`**.

- **URL:** `/api/v1/generate-key`
- **Method:** `GET`
- **Functionality:** Generates and returns a new, cryptographically secure short key that can be used by the `url-shortener-service`.

