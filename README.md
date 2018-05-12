# go-rest-api-template

Template for building REST Web Services in Golang with:

* [go-swagger](https://github.com/go-swagger/go-swagger) for representation of RESTful API
* [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) for routing
* [codegangsta/negroni](https://github.com/codegangsta/negroni) as a middleware handler
* [strechr/testify](https://github.com/stretchr/testify) for writing easier test assertions
* [unrolled/render](https://github.com/unrolled/render) for HTTP response rendering
* [palantir/stacktrace](https://github.com/palantir/stacktrace) to provide more context to error messages
* [unrolled/secure](https://github.com/unrolled/secure) to improve API security
* [oxeque/realize](https://github.com/oxequa/realize) to watch for changes in files during development

## API Routes with Health check

```
Route{"Healthcheck", "GET", "/healthcheck", HealthcheckHandler},
//=== USERS ===
Route{"ListUsers",  "GET", "/users", ListUsersHandler},
Route{"GetUser",    "GET", "/users/{uid:[0-9]+}", GetUserHandler},
Route{"CreateUser", "POST", "/users", CreateUserHandler},
Route{"UpdateUser", "PUT", "/users/{uid:[0-9]+}", UpdateUserHandler},
Route{"DeleteUser", "DELETE", "/users/{uid:[0-9]+}", DeleteUserHandler},
```

## API specification

Execute the following in the directory with the main.go and it will parse all the files that are reachable by that main package to produce a swagger specification and serve it with SwaggerUI:

```
swagger generate spec -o ../../swagger.json --scan-models
swagger serve -F=swagger ../../swagger.json
```

## Testing

Retrieve a list of users:

```
curl -X GET http://localhost:3009/users | jq
```

Get a specific user:

```
curl -X GET http://localhost:3009/users/0 | jq
```

## Starting the service on a production server

This is how you could run your app on a server:

First, you copy the binary and the `fixtures.json` + `VERSION` files into a directory, e.g. `/opt/go-rest-api-template`.

Then start the app as a service. Store the app's PID in a text file so we can kill it later.

```
#!/bin/bash
export ENV=DEV
export PORT=8080
export VERSION=/opt/go-rest-api-template/VERSION
export FIXTURES=/opt/go-rest-api-template/fixtures.json
sudo nohup /opt/go-rest-api-template/go-rest-api-template >> /var/log/go-rest-api-template.log 2>&1&
echo $! > /opt/go-rest-api-template/go-rest-api-template-pid.txt
```

When you want to kill your app later during a redeployment or a server shutdown, then you can kill the app by looking up the previously stored PID:

```
#!/bin/bash
if [ -f /opt/go-rest-api-template/go-rest-api-template-pid.txt ]; then
  kill -9 `cat /opt/go-rest-api-template/go-rest-api-template-pid.txt`
  rm -f /opt/go-rest-api-template/go-rest-api-template-pid.txt
fi
```
