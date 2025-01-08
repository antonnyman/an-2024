
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
EXPOSE 5000
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]

