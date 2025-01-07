# Multi-stage build for cross-platform compatibility
FROM golang:1.22.7-alpine AS builder

RUN apk add --no-cache git openssh-client

# Set up environment variables for cross-compilation
ARG TARGETOS
ARG TARGETARCH
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

WORKDIR /app

# Copy Go module files and download dependencies
COPY ./src/go.mod ./src/go.sum ./
COPY ./src/vendor ./
#RUN go mod download

# Copy application source code
COPY ./src .

# Build the binary for the target platform
RUN go build -o authentication-service

# Final lightweight image
FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup


WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/authentication-service .

# Change ownership of the application directory
RUN chown -R appuser:appgroup /app

# Switch to the non-root user
USER appuser

# Expose the application port
EXPOSE 2000

# Run the application
CMD ["./authentication-service"]