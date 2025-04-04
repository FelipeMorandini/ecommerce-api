FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ecommerce-api ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/ecommerce-api .

# Copy the .env file if it exists (for development)
COPY --from=builder /app/.env* ./ 2>/dev/null || true

EXPOSE 3000

CMD ["./ecommerce-api"]
