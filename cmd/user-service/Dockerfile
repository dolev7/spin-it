FROM golang:1.21 AS builder

# Set the working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o user-service cmd/user-service/main.go

# Use a minimal base image for running the app
FROM gcr.io/distroless/base-debian11

# Set the working directory
WORKDIR /app

# Copy the compiled binary from builder
COPY --from=builder /app/user-service .

# Expose the port
EXPOSE 8080

# Run the binary
CMD ["/app/user-service"]
