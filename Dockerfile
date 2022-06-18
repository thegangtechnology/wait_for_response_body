# Build the app
FROM golang:1.13 as builder
WORKDIR /app
# Enable go modules even inside GOPATH
ENV GO111MODULE=on

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o wait_for_response /app/main/main.go

# Create a minimal docker container and copy the app into it
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/wait_for_response .
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
ENTRYPOINT ["/app/entrypoint.sh"]