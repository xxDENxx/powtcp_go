FROM golang:1.22-bookworm AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o client ./main.go

FROM debian:bookworm-slim

COPY --from=builder /build/client /

ENTRYPOINT ["/client"]
