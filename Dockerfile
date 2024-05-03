# Build the app
FROM golang:1.22 as builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64=value
WORKDIR /app
COPY ./ ./
RUN go mod download && \
    go test ./... && \
    go build -a -installsuffix cgo -o wait_for_response /app/main/main.go

# Create a minimal docker container and copy the app into it
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/wait_for_response .
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
ENTRYPOINT ["/app/entrypoint.sh"]
