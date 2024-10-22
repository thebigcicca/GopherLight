# Build Stage
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files from the root of the project
COPY ../../go.mod ../../go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY ../../router ./router
COPY ../../req ./req
COPY main.go ./

# Build the Go app
RUN go build -o crud_app

# Run Stage
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/crud_app ./

# Expose port 3333 to the outside world
EXPOSE 3333

# Command to run the executable
CMD ["./crud_app"]
