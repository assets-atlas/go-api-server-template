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