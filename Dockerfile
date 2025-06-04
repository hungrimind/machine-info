# Start with the official Golang image to build the app
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code
COPY main.go .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o linux-version-service

# Final minimal image
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/linux-version-service /linux-version-service

# Expose the port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/linux-version-service"]
