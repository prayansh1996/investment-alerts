# Start from the official Go image to build your application
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o alerts .

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/alerts .
COPY --from=builder /app/holdings/holdings.yaml .

# Expose port 9090 to the outside world
EXPOSE 9090

# Command to run the executable
CMD ["./alerts"]
