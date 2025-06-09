FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .

RUN chmod +x ./main

EXPOSE 8080
CMD ["./main"]
# This Dockerfile builds a Go application and prepares it for deployment in a minimal Alpine Linux container.