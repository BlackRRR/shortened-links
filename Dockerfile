# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:latest

WORKDIR /app

COPY . .
RUN go build ./cmd/shortened-links

CMD ["./shortened-links"]
