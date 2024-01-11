# Stage 1: Build Stage
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the Go application source code and Protobuf files to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/cmd/main ./cmd/main.go

# Stage 2: Final Stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary artifacts from the previous stage
COPY --from=builder /app/cmd/main ./cmd/main
COPY --from=builder /app/conf ./conf

# Expose the port that your gRPC server will run on
EXPOSE 50050

# Command to run your application
CMD ["./cmd/main"]
