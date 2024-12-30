# Stage 1: Install dependencies
FROM golang:1.23.1-bookworm AS deps

WORKDIR /app

# Copy go.mod and go.sum to the container
COPY go.mod go.sum ./

# Run `go mod download` to download all dependencies
RUN go mod download

# Verify and debug if necessary - check if the module is correctly fetched
RUN go list -m all
RUN ls /go/pkg/mod/github.com/!serhii!cho/timeago@v0.0.0-20231226174358-3bade6b97419/langs

# Stage 2: Build the application
FROM golang:1.23.1-bookworm AS builder

WORKDIR /app

# Copy the Go module cache from the deps stage
COPY --from=deps /go/pkg/mod /go/pkg/mod

# Copy the rest of the application files
COPY . .

# Enable them if you need them
ENV CGO_ENABLED=0
ENV GOOS=linux

# Build the Go application
RUN go build -ldflags="-w -s" -o main cmd/main.go

# Final stage: Run the application
FROM debian:bookworm-slim

WORKDIR /app

# Create a non-root user and group
RUN groupadd -r appuser && useradd -r -g appuser appuser

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Change ownership of the application binary
RUN chown appuser:appuser /app/main

# Switch to the non-root user
USER appuser

# Set the default command to run the application
CMD ["./main"]
