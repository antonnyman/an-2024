# Build Stage
FROM golang:1.21-alpine as builder

WORKDIR /app

# Install any dependencies for templ, if needed
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go fmt && \
    templ fmt ./views && \
    templ generate && \
    go build -o /app/main .

# Production Stage with Alpine
FROM alpine:latest

WORKDIR /app

# Copy only the compiled binary
COPY --from=builder /app/main /app/main

# Expose the required port
EXPOSE 8000

# Run the application
CMD ["/app/main"]
