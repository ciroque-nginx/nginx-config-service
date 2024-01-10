# Developing

This guide helps get up and running to contribute code support to the project.

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

### Build the Docker Image

```bash
docker build -t ciroque-nginx/nginx-config-service:latest .
```
