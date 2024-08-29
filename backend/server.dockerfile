
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

WORKDIR /app/server

RUN go mod tidy

RUN go build -o /app/server-build ./cmd/api/

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server-build .

CMD ["./server-build"]
