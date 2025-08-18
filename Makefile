# Define variables for common paths and commands
REDIRECT_SERVICE_PATH := ./url-redirect-service
SHORTEN_SERVICE_PATH := ./url-shortener-service
KEYGEN_SERVICE_PATH := ./key-gen-service
GO_CMD := go

# Define flags for specific services and test types
# r = redirect service
# s = shorten service
# k = keygen service
SERVICES ?= rsk
TYPE ?= all

.PHONY: test check-redis check-postgres clean clean-redirect clean-shorten clean-keygen

# =================================================================
# Main Test Target
# =================================================================

# test: This is the main entry point for running tests.
# It uses the SERVICES and TYPE variables to determine
# which tests to run.
#
# Examples:
#
#   make test                   (Runs all tests for all services)
#
#   make test SERVICES=r   (Runs all tests for only the Redirect service)
#
#   make test SERVICES=rs  (Runs all tests for the Redirect and Shorten services)
#
#   make test SERVICES=r TYPE=api
#                               (Runs only API tests for the Redirect service)
#
#   make test SERVICES=s TYPE=services
#                               (Runs only core services tests for the Shorten service)
#
#   make test TYPE=storage (Runs storage tests for all services that have them)
test:
	@echo "--- Running tests for services: $(SERVICES) ---"
	@if [ "$(TYPE)" = "all" ]; then \
		echo "Running all tests..."; \
	elif [ "$(TYPE)" = "api" ]; then \
		echo "Running API tests..."; \
	elif [ "$(TYPE)" = "services" ]; then \
		echo "Running core services tests..."; \
	elif [ "$(TYPE)" = "storage" ]; then \
		echo "Running storage tests..."; \
	elif [ "$(TYPE)" = "web" ]; then \
		echo "Running web tests..."; \
	else \
		echo "Error: Invalid TYPE. Use 'all', 'api', 'services', 'web', or 'storage'."; \
		exit 1; \
	fi

	@if echo "$(SERVICES)" | grep -q "r"; then \
		$(MAKE) test-redirect-$(TYPE); \
	fi
	@if echo "$(SERVICES)" | grep -q "s"; then \
		$(MAKE) test-shorten-$(TYPE); \
	fi
	@if echo "$(SERVICES)" | grep -q "k"; then \
		$(MAKE) test-keygen-$(TYPE); \
	fi

# =================================================================
# Service-Specific Test Targets
# =================================================================

.PHONY: test-redirect-all test-redirect-api test-redirect-services test-redirect-storage test-redirect-web
test-redirect-all: check-redis
	@echo "--- Running all tests for the URL Redirect Service ---"
	$(GO_CMD) test -v -cover $(REDIRECT_SERVICE_PATH)/...
test-redirect-api: check-redis
	@echo "--- Running API tests for the URL Redirect Service ---"
	$(GO_CMD) test -v -cover $(REDIRECT_SERVICE_PATH)/internal/api/...
test-redirect-services:
	@echo "--- Running core services tests for the URL Redirect Service ---"
	$(GO_CMD) test -v -cover $(REDIRECT_SERVICE_PATH)/internal/core/services/...
test-redirect-storage: check-redis
	@echo "--- Running storage tests for the URL Redirect Service ---"
	$(GO_CMD) test -v -cover $(REDIRECT_SERVICE_PATH)/internal/infrastructure/storage/...
test-redirect-web: check-redis
	@echo "--- Running web tests for the URL Redirect Service ---"
	$(GO_CMD) test -v -cover $(REDIRECT_SERVICE_PATH)/internal/infrastructure/web/...

.PHONY: test-shorten-all test-shorten-api test-shorten-services test-shorten-storage test-shorten-web
test-shorten-all: check-postgres
	@echo "--- Running all tests for the URL Shortener Service ---"
	$(GO_CMD) test -v -cover $(SHORTEN_SERVICE_PATH)/...
test-shorten-api: check-postgres
	@echo "--- Running API tests for the URL Shortener Service ---"
	$(GO_CMD) test -v -cover $(SHORTEN_SERVICE_PATH)/internal/api/...
test-shorten-services:
	@echo "--- Running core services tests for the URL Shortener Service ---"
	$(GO_CMD) test -v -cover $(SHORTEN_SERVICE_PATH)/internal/core/services/...
test-shorten-storage: check-postgres
	@echo "--- Running storage tests for the URL Shortener Service ---"
	$(GO_CMD) test -v -cover $(SHORTEN_SERVICE_PATH)/internal/infrastructure/storage/...
test-shorten-web: check-postgres
	@echo "--- Running web tests for the URL Shortener Service ---"
	$(GO_CMD) test -v -cover $(SHORTEN_SERVICE_PATH)/internal/infrastructure/web/...


.PHONY: test-keygen-all test-keygen-services test-keygen-api
test-keygen-all:
	@echo "--- Running all tests for the Key Generation Service ---"
	$(GO_CMD) test -v -cover $(KEYGEN_SERVICE_PATH)/...
test-keygen-services:
	@echo "--- Running core services tests for the Key Generation Service ---"
	$(GO_CMD) test -v -cover $(KEYGEN_SERVICE_PATH)/internal/core/services/...
test-keygen-api:
	@echo "--- Running API tests for the Key Generation Service ---"
	$(GO_CMD) test -v -cover $(KEYGEN_SERVICE_PATH)/internal/api/...

# =================================================================
# Helper Targets
# =================================================================

# check-redis: Ensures the Redis container is running before running tests.
check-redis:
	@echo "--- Checking for running Redis container ---"
	@if [ -z "$$(docker ps -q -f name=my-redis)" ]; then \
		echo "Redis container not found. Starting it..."; \
		docker run --name my-redis -p 6379:6379 -d --rm redis; \
	fi

check-postgres:
# TODO: implement this once finished with shortener service

# clean: Cleans up Go test cache and build artifacts for all services.
clean:
	@echo "--- Cleaning Go test cache ---"
	$(GO_CMD) clean -testcache
	@echo "--- Tidying modules for all services ---"
	@cd $(REDIRECT_SERVICE_PATH) && $(GO_CMD) mod tidy
	@cd $(SHORTEN_SERVICE_PATH) && $(GO_CMD) mod tidy
	@cd $(KEYGEN_SERVICE_PATH) && $(GO_CMD) mod tidy

clean-redirect:
	@echo "--- Cleaning Go test cache ---"
	$(GO_CMD) clean -testcache
	@echo "--- Tidying modules for Redirect service ---"
	@cd $(REDIRECT_SERVICE_PATH) && $(GO_CMD) mod tidy

clean-shorten:
	@echo "--- Cleaning Go test cache ---"
	$(GO_CMD) clean -testcache
	@echo "--- Tidying modules for Shortener service ---"
	@cd $(SHORTEN_SERVICE_PATH) && $(GO_CMD) mod tidy

clean-keygen:
	@echo "--- Cleaning Go test cache ---"
	$(GO_CMD) clean -testcache
	@echo "--- Tidying modules for Keygen service ---"
	@cd $(KEYGEN_SERVICE_PATH) && $(GO_CMD) mod tidy
