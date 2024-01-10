# Copyright Steve Wagner. All rights reserved.
# Use of this source code is governed by the Apache License that can be found in the LICENSE file.

FROM golang@sha256:2523a6f68a0f515fe251aad40b18545155135ca6a5b2e61da8254df9153e3648 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o nginx-config-service ./cmd/main/main.go

FROM alpine@sha256:ca452b8ab373e6de9c39d030870a52b8f0d3a9cf998c43183fd114660ae96330

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
