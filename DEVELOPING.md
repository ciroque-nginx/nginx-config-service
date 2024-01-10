# Developing

This guide helps get up and running to contribute code to the project.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.21 or later
- [Docker](https://docs.docker.com/get-docker/) 20.10.6 or later

If you use asdf, you can install the prerequisites using the following commands:

```bash
asdf plugin add golang
asdf install 
```

## Getting Started

### Clone the Repository

```bash
git clone git@github.com:ciroque-nginx/nginx-config-service.git
```

### Run the Tests

```bash
go test ./...
```

### Build the Project

```bash
go build -o _build/ncs ./cmd/main/main.go
```

### Building the Container

Configuration options can be specified using `--build-args` when building the container. For example, to specify the port that NCS listens on, use the `PORT` build argument.

```bash
docker build --build-arg NCS_HTTP_PORT=8080 -t nginx-config-service .
```

The available build arguments are:
- NCS_HTTP_PORT, the port that NCS listens on, defaults to 8293.
- NCS_LOG_LEVEL, the log level for NCS, defaults to INFO.
- NCS_HTTP_HOST, the host that NCS listens on, defaults to 0.0.0.0.
