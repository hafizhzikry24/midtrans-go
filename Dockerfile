# Use official Golang image to build the Go app
FROM golang:1.23-alpine as builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main .

# Start a new stage from a smaller image
FROM alpine:latest

# Install necessary dependencies (for example, ca-certificates)
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app is running on (e.g., 8080)
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
