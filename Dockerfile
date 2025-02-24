# Use the official Go image as the base image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install dependencies for bcrypt hashing
RUN apk add --no-cache git gcc musl-dev

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o library-management ./cmd/main.go

# Build the seeder
RUN CGO_ENABLED=0 GOOS=linux go build -o seeder ./scripts/seed.go

#  Final lightweight image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/library-management .
COPY --from=builder /app/seeder .

# Install Go runtime for running the seeder script
RUN apk add --no-cache go libc6-compat

# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
CMD ["./library-management"]