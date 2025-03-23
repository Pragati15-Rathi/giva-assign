#Use the official Golang image as a builder
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o giva 

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/giva .

EXPOSE 3000

CMD ["./giva"]

