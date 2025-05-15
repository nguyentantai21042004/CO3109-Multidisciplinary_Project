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

# Build ứng dụng
RUN swag init -g cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

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
COPY --from=builder /app/api .

# Ensure the binary is executable
RUN chmod +x /app/api

# Expose port 8085
EXPOSE 8085

# Run the application
ENTRYPOINT ["/app/api"] 