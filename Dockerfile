# Base image with Go
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Install required system packages
RUN apk add --no-cache gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o progserver ./cmd/server

# Expose gRPC port
EXPOSE 8080

# Run the server
CMD ["./progserver"]