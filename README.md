[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![Community Support](https://badgen.net/badge/support/community/cyan?icon=awesome)](https://github.com/nginxinc/ciroque-nginx/nginx-config-service/blob/main/SUPPORT.md)

[![Project Status: Concept – Minimal or no implementation has been done yet, or the repository is only intended to be a limited example, demo, or proof-of-concept.](https://www.repostatus.org/badges/latest/concept.svg)](https://www.repostatus.org/#concept)

# NGINX Config Service

The NGINX Config Service, or NCS, is a service that provides a RESTful API for managing NGINX configuration files. The NCS is a microservice that is intended to be deployed as a sidecar to an NGINX instance. The NCS is written in Go and is designed to be deployed as a container.

## Table of Contents

- [Overview](#overview)
- [Goals](#goals)
- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [How to Use](#how-to-use)
- [Contributing](#contributing)
- [License](#license)

## Overview

[NGINX](https://www.nginx.com/) is a high-performance, high-concurrency web server that is used by some of the largest websites in the world. NGINX is also used as a reverse proxy, load balancer, and API gateway. NGINX is a very flexible and powerful tool that can be used to solve a wide variety of problems.
One of the weak points of NGINX is dynamic configuration. It is possible to dynamically configure NGINX using the [NGINX Plus API](https://docs.nginx.com/nginx/admin-guide/load-balancer/http-api/), but this requires NGINX Plus; there are other solutions as well, see [nginx agent](https://github.com/nginx/agent).

This project aims to simplify the management of NGINX configuration, specifically for NGINX Open Source, though it should work with NGINX Plus as well.

## Goals

- Provide a RESTful API for managing NGINX configuration files.
- Allow dynamic configuration of NGINX without requiring NGINX Plus.
- Provide a simple, easy to use, and easy to understand API.
- Support full NGINX configuration syntax.

## Non-Goals

- Provide monitoring or metrics from NGINX instance.
- Provide fleet management of NGINX instances.
- Provide a GUI for managing NGINX configuration files.

## Requirements

- NGINX 1.19.0 or later.

## Getting Started

It depends. What are your goals?

- To run NCS and manage NGINX configuration files; see [How to Use](#how-to-use).
- To contribute to the NCS project; see [Contributing](#contributing).

## How to Use

The easiest way to use NCS is to run it as a container. The following example shows how to run NCS as a container using Docker.

```bash
docker run -dit --name nginx-config-service -p 8293:8293 ghcr.io/ciroque/nginx-config-service:latest
```

Note that using `latest` is not recommended for production use. Instead, use a specific version tag, such as `v0.1.0`.

The project provides [releases](https://github.com/ciroque-nginx/nginx-config-service/releases) to help find a specific version.

## Contributing

Please see the [contributing guide](https://github.com/ciroque-nginx/nginx-config-service/blob/main/CONTRIBUTING.md) for guidelines on how to best contribute to this project.

Also, see the [developer guide](https://github.com/ciroque-nginx/nginx-config-service/blob/main/DEVELOPING.md) for information on how to build and test this project.

## License

[Apache License, Version 2.0](https://github.com/ciroque-nginx/nginx-config-service/blob/main/LICENSE)

