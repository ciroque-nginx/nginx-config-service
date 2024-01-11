# Design

## Overview

### Architecture

```mermaid
```

### Components

The following components are used in the system:

- **Server**: Standard HTTP server, handles incoming requests and passes them to the handler.
- **HttpHandler**: Handles incoming requests, passes them to the request processor and returns the response.
- **RequestProcessor**: Processes incoming requests, returns the response.
- **Updater**: Updates the NGINX config.

#### Server

The server is a standard HTTP server, which handles incoming requests and passes them to the handler. Routes
are defined by handler implementations; the list of Routes is defined in the `handler.go` file.

#### HttpHandler

The HTTP handler handles incoming requests by creates a concrete `configs.Interface` instance and passing it to the update processor.
An HTTP 202 Created response that includes a Location header containing the URL of the created resource and JSON body containing the
UpdateEvent detail is returned.

Handler implementation defines a route. The planned routes are:

- `/nginx/config/raw/{name}`: Create / Delete raw NGINX config
- `/nginx/config/upstream/{name}`: Create / Delete an upstream config
- `/nginx/config/server/{name}`: Create / Delete a server config

Note: only POST and DELETE methods are supported.

#### ConfigUpdateProcessor

Given a concrete `configs.Interface` instance, the update processor creates a new `UpdateEvent` and passes it to the updater via channel.

#### UpdateEvent

The `UpdateEvent` contains the information needed to build the NGINX config. It also maintains a log of events that occur during the
update process.

An `UpdateEvent` is initialized in the `Accepted` state. Each step along the way, the `UpdateEvent` is updated with the current state.

If, at any step, the update fails, the `UpdateEvent` is updated with the current state and the `Error` field is set to the error message.

#### UpdateProcessor

- Maintains a buffered channel of `UpdateEvent` instances
- Updates the NGINX config by saving a file in the NGINX config directory
- Filename is provided as the final route segment
- Once file is written, invoke `nginx -t` to test the config; if successful, invoke `nginx -s reload` to reload the config.
 