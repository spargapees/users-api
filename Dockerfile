FROM golang:latest

WORKDIR /app

# Check Go version
RUN go version

# Set environment variables
ENV GOPATH=/

# Copy the application code
COPY ./ ./

# Install PostgreSQL client (optional if you need it for scripts)
RUN apt-get update && apt-get -y install postgresql-client

# Ensure the wait script is executable if used
RUN chmod +x wait-for-postgres.sh

# Download Go module dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o users-api ./main.go

# Command to start the application
CMD ["./users-api"]
