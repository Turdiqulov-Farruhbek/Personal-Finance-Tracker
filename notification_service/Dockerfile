FROM golang:1.22.1 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o notification ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/notification .
COPY --from=builder /app/helper/format.html ./helper/format.html


COPY .env .
EXPOSE 30030
CMD ["./notification"]