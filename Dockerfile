# Description: Dockerfile for building the auth-service image
FROM golang:1.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o auth-service cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/auth-service /usr/bin/auth-service

# Command to run the executable
RUN ls -l /usr/bin/auth-service

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/usr/bin/auth-service"]