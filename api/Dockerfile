# Start from the Go image for building the API
FROM golang:1.20.13 as builder

# Set the working directory in the container
WORKDIR /app

# Copy go mod and sum files
COPY go.* ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY ./ ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage from scratch for a smaller final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
