FROM golang:1.16.3-alpine AS builder
RUN mkdir /kafkawebinar
ADD . /kafkawebinar
WORKDIR /kafkawebinar

RUN apk add --no-cache git
RUN go build -o bin/consumer -ldflags="-s -w" ./cmd/consumer && \
    go build -o bin/producer -ldflags="-s -w" ./cmd/producer

FROM alpine:latest
RUN mkdir -p /go/bin && \
    mkdir /logs && \
    chmod -R 755 /logs
WORKDIR /go/bin
COPY --from=builder /kafkawebinar/bin/consumer .
COPY --from=builder /kafkawebinar/bin/producer .