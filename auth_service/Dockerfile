FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/auth .

COPY --from=builder /app/helper/format.html ./helper/

COPY .env .env 

RUN chmod +x auth

EXPOSE 8080

CMD ["./auth"]