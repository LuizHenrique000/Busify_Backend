FROM golang:1.23-bookworm AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /run-app ./cmd/server

FROM debian:bookworm

WORKDIR /app
COPY --from=builder /run-app /run-app

EXPOSE 8080
CMD ["/run-app"]
