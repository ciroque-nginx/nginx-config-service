# Copyright Steve Wagner. All rights reserved.
# Use of this source code is governed by the Apache License that can be found in the LICENSE file.

FROM golang:1.21.6-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o nginx-config-service ./cmd/main/main.go

FROM alpine:3.16

ARG NCS_LOG_LEVEL="info"
ARG NCS_HTTP_HOST="0.0.0.0"
ARG NCS_HTTP_PORT="8293"

ENV NCS_LOG_LEVEL=${NCS_LOG_LEVEL} \
    NCS_HTTP_HOST=${NCS_HTTP_HOST} \
    NCS_HTTP_PORT=${NCS_HTTP_PORT}

WORKDIR /app

RUN adduser -u 11115 -D -H  ncs

USER ncs

COPY --from=builder /app/nginx-config-service .

ENTRYPOINT ["/app/nginx-config-service"]
