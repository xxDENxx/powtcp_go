FROM golang:1.22-bookworm AS builder

WORKDIR /build
COPY . /build
RUN go mod download
RUN go build -o server ./main.go

FROM debian:bookworm-slim

COPY --from=builder /build/server /var/www/server

ENTRYPOINT ["/var/www/server"]
