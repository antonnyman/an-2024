# Fetch
FROM golang:latest AS fetch-stage
COPY go.mod go.sum /app
WORKDIR /app
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build
FROM golang:latest AS build-stage
COPY --from=generate-stage /app /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app

# Test
FROM build-stage AS test-stage
RUN go test -v ./...

# Deploy
FROM gcr.io/distroless/base-debian12 AS deploy-stage
WORKDIR /
COPY --from=build-stage /app/app /app
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app"]

## Install any dependencies for templ, if needed
#RUN apk add --no-cache git
#
#RUN go mod download
#
#COPY . .
#
#RUN go fmt && \
#    templ fmt ./views && \
#    templ generate && \
#    go build -o /app/main .
#
## Production Stage with Alpine
#FROM alpine:latest
#
#WORKDIR /app
#
## Copy only the compiled binary
#COPY --from=builder /app/main /app/main
#
## Expose the required port
#EXPOSE 8000
#
## Run the application
#CMD ["/app/main"]
