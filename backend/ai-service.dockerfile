
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

WORKDIR /app/ai-service

RUN go mod tidy

RUN go build -o /app/ai-service-build ./cmd/api/

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/ai-service-build .

CMD ["./ai-service-build"]
