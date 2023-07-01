# Golang API Server Template Repo

This template repo can be used as a scaffolding project to create new API servers.

## Server Specification

- The server runs on port 2000 by default. This can be overridden by setting the `SERVER_HTTP_PORT` environment variable
- There is a default route enabled at the `/` endpoint. This responds only to `GET` calls and returns the service name and version number
- The service name by default is `example-service`. This can be overridden by setting the `SERVER_SERVICE_NAME` environment variable
- The version by default is set to `0.1.0`. This can be overridden by setting the `SERVER_VERSION` environment variable.
- The server is built using `net/http` and `mux`.

## Routes

All routes can be added to the `NewRouter()` function in the `server.go` file.

**Example Route**

```go
r.HandleFunc("/signup", SignUp(db)).Methods("POST")
```

## Logging

There are extensive logging capabilities built into this template. There are 7 levels of logging:

- error
- warn
- info
- debug
- trace
- fatal
- panic

By default, log outputs will be set at the `info` level; however, this can be overridden by setting the `LOG_LEVEL` environment variable.

Logs are output in JSON format to assist with log ingestion into 3rd party system, i.e Splunk.

To add logging to more functions, use `log` followed by the log level.

```go
log.Warn(err)
```

Logging can also be more bespoke. For example, the server startup logs contains custom schema

```go
log.WithFields(
    log.Fields{
        "server_startup_info": log.Fields{
            "service_name": serviceName,
            "http_port":    httpPort,
            "version":      version,
            "log_level":    logLevel,
            "hostname":     hostName,
        }},
).Info("server started...")
```

Which renders the following log output:

```json
{
  "level": "info",
  "msg": "server started...",
  "server_startup_info": {
    "hostname": "roberts-mac-pro.lan",
    "http_port": "2000",
    "log_level": "debug",
    "service_name": "example-service",
    "version": "0.1"
  },
  "time": "2023-07-01T01:04:50+01:00"
}
```
## Dependencies

You will need to initialize a new go.mod file with your desired name in the `./src` folder:

```shell
go mod init <Name of your package>
```

Once this is done, you can run:

```shell
go mod tidy
```

This will download all of your package dependencies.

## Makefile

There is a Makefile setup which allows for building and test locally, as well docker building, tagging and pushing of your microservice.

In order to use the Makefile, update the variables at the top of the make file as per your project:

```makefile
service_name = "example-service"
version = "0.1.0"
org = "assets-atlas"

```