FROM golang:1.21-alpine AS builder

RUN apk add curl gcc g++
RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN templ generate
RUN GOOS=linux CGO_ENABLED=1 GARCH=$(arch) go build -o ./.build/main ./cmd

FROM alpine:3.19.0 AS runtime

LABEL org.opencontainers.image.source=https://github.com/maantje/postcode

RUN apk update && apk add gzip

COPY --from=builder /app/.build/main /usr/local/bin/postcode
COPY --from=ghcr.io/maantje/postcode-database:latest /database/addresses.sqlite.gz /application/database/addresses.sqlite.gz

WORKDIR /application

COPY start.sh start.sh

EXPOSE 80

CMD ["./start.sh"]
