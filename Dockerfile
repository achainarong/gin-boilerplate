# syntax=docker/dockerfile:1

FROM golang:1.20.3-alpine3.17 AS builder

RUN apk update && apk add --no-cache 'git=~2'
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

FROM alpine:3.17

WORKDIR /

COPY --from=builder /app/main /app/main

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

WORKDIR /app

ENTRYPOINT ["/app/main"]
