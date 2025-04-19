# Base image 
FROM golang:1.24-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /code

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# ? CGO_ENABLED=0 disables cgo, which is necessary for a statically linked binary
# ? GOOS=linux sets the target operating system to Linux
# ? -a forces rebuilding of packages that are already up-to-date
# ? -installsuffix cgo appends a suffix to the package name to avoid conflicts with cgo-enabled packages
# ? -o main specifies the output binary name
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Start a new stage from scratch
FROM alpine:3.21

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /code/main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Run as non-root user
RUN adduser -D myuser

USER myuser

# Command to run the executable
CMD ["./main"]