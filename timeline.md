### Project Timeline: Microservices Architecture ⏳

This project timeline is a sequential breakdown of the development process for your microservices-based URL shortener. Each phase builds upon the last, ensuring a logical and coherent development flow.

***

### Phase 1: Foundation & Shared Components (Days 1-3) 🏗️

This initial phase focuses on establishing the project's foundational structure, shared resources, and configuration.

- **Set up the Monorepo:** Create the `duss/` root directory and all service subdirectories (`api-gateway-service`, `key-gen-service`, etc.).
- **Define Shared Domain Model:** Create `shared/domain/url.go` to define the `URL` data structure. This is crucial for consistent data handling across all services.
- **Configure Docker and Orchestration:** Create the initial `docker-compose.yml` file to define the services (PostgreSQL, Redis, and the Go services). Start with a simple setup to verify services can communicate.
- **Establish Project-level Files:** Create the `Makefile`, `README.md`, `LICENSE`, and other essential documentation files to ensure good project hygiene from the start.

***

### Phase 2: Core Service Development (Days 4-10) 💻

This is the main development phase where you build the isolated logic for each of the core backend services. You will focus on their internal functionality and testing.

- **Develop the `key-gen-service`:**
    - Create the `internal/core/services/key_generator.go` with the core logic for generating unique keys.
    - Write the internal API handlers in `internal/api/handlers.go` and the router in `internal/infrastructure/web/router.go`.
    - Implement `key_generator_test.go` and `handlers_test.go` to ensure correctness.
- **Develop the `url-shortener-service`:**
    - Implement the core shortening logic in `internal/core/services/shortener.go`.
    - Create the storage layer (`internal/infrastructure/storage`) with implementations for PostgreSQL and Redis (`postgres.go`, `redis.go`).
    - Write `shortener_test.go` and the storage tests (`postgres_test.go`, `redis_test.go`).
    - **Crucially, develop the `url-shortener-service` to call the `key-gen-service`'s internal API to get a key.**
- **Develop the `url-redirect-service`:**
    - Implement the core redirection logic in `internal/core/services/redirect.go`.
    - Create the storage layer (`internal/infrastructure/storage`) for retrieving the original URL.
    - Write `redirect_test.go` and the storage tests to ensure the service can correctly fetch URLs.

***

### Phase 3: API Gateway & Integration (Days 11-14) 🌐

This phase focuses on building the public-facing API Gateway and integrating all services.

- **Develop the `api-gateway-service`:**
    - Implement the core orchestration logic in `internal/core/services/api_gateway.go`. This service should not have business logic; its purpose is to route requests.
    - Create the internal client implementations (`internal/infrastructure/clients/shortener_client.go`, `redirect_client.go`) that communicate with the backend services.
    - Define the public-facing API endpoints in `internal/infrastructure/web/router.go`.
- **Implement API Handlers:** Write the public handlers in `internal/api/handlers.go` for `/shorten` and `/:shortKey`.
    - The `/shorten` handler should call the `url-shortener-service` via its client.
    - The `/:shortKey` handler should call the `url-redirect-service` via its client and issue a 301 redirect.
- **Integrate Services:** Update the `docker-compose.yml` to define the network and services, ensuring the API Gateway is the only service with exposed ports. Verify all inter-service communication works correctly within the Docker network.
- **Write Integration Tests:** Create comprehensive tests for the `api-gateway-service` to ensure it correctly delegates requests and handles responses from the backend services.

***

### Phase 4: Finalization & Deployment (Days 15-18) ✅

The final phase involves polishing the project, creating deployment scripts, and preparing for a release.

- **Add Final `Dockerfile`s:** Create `Dockerfile`s for each service (`url-shortener-service`, `url-redirect-service`, etc.) to create self-contained images.
- **Refine Configuration:** Finalize the `configs/config.yaml` file to include all necessary environment variables for each service, such as database credentials and internal API URLs.
- **Scripting:** Create the `scripts/init.sh` script to automate project setup, such as running the `docker-compose` stack and performing any necessary database migrations.
- **Documentation:**
    - Complete the `README.md` with detailed instructions on how to set up and run the project.
    - Write the `ARCHITECTURE.md` file to clearly explain the microservices design and data flow.
    - Finalize the `CONTRIBUTING.md` and `CODE_OF_CONDUCT.md`.
- **Final Review:** Perform a final code review and run all tests to ensure the entire system is stable and all components are working as expected.
