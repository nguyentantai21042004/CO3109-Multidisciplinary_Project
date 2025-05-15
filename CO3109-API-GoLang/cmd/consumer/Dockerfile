##
## STEP 1 - BUILD APPLICATION AND CREATE BINARY FILE
##

FROM golang:1.23.8-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.7

# Copy source code
COPY . .

# Build application
RUN swag init -g cmd/consumer/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o consumer ./cmd/consumer

##
## STEP 2 - COPY BINARY FILE FROM BUILD STAGE TO A LIGHTER IMAGE (ALPINE) TO RUN APPLICATION
##
FROM alpine:latest

# Set timezone
ENV TZ=Asia/Ho_Chi_Minh

# Install necessary packages and set timezone
RUN apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

# Copy the binary from the builder stage
WORKDIR /app
COPY --from=builder /app/consumer .

# Ensure the binary is executable
RUN chmod +x /app/consumer

# Expose port 80
EXPOSE 80

# Run the application
ENTRYPOINT ["/app/consumer"]