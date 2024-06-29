FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy go mod and sum files
COPY go.* ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory
COPY main.go .
COPY handler.go .

# Build the Go app
RUN go build -o main .

# Run the Go app
FROM golang:1.22-alpine AS runner

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /usr/src/app/main /usr/src/app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/usr/src/app/main"]
