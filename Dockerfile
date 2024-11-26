FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/main .

RUN apt-get update && apt-get install -y libc6

EXPOSE 8080

CMD ["./main"]
