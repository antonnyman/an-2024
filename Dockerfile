
# Fetch
FROM golang:latest AS fetch-stage
ENV LANG=C.UTF-8 LC_ALL=C.UTF-8
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
ENV LANG=C.UTF-8 LC_ALL=C.UTF-8
WORKDIR /app
COPY --chown=65532:65532 . /app
RUN ["templ", "generate"]

# Build
FROM golang:latest AS build-stage
ENV LANG=C.UTF-8 LC_ALL=C.UTF-8
WORKDIR /app
COPY --from=generate-stage /app /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app

# Deploy
FROM gcr.io/distroless/base-debian12 AS deploy-stage
ENV LANG=C.UTF-8 LC_ALL=C.UTF-8
WORKDIR /app
COPY --from=build-stage /app/app /app/app
COPY --from=build-stage /app/assets /app/assets
#RUN chmod +x /app/app  # Ensure the binary has executable permissions
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]


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
