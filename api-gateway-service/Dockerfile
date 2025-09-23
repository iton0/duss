# Stage 1: The Build Stage
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy all the project files from the root duss directory
COPY . .

# Change the working directory to the API Gateway service directory
WORKDIR /app/api-gateway-service

# Build the Go application binary. The -o flag names the output binary.
# The `.../cmd/server` is the main package to build.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /api-gateway ./cmd/server

# ---

# Stage 2: The Final, Minimal Image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /api-gateway .

# Expose the port your application listens on (e.g., 8080)
EXPOSE 8080

# Create a non-root user and switch to it
RUN addgroup -g 1001 gatewaygroup && \
    adduser -u 1001 -G gatewaygroup -D gatewayuser
USER gatewayuser

# Command to run the executable
CMD ["./api-gateway"]
