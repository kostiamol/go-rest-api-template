# go-rest-api-template

Template for building REST Web Services in Golang. Uses gorilla/mux as a router/dispatcher and Negroni as a middleware handler.

## Introduction

### Why?

After writing many REST APIs with Java Dropwizard, Node.js/Express and Go, I wanted to distill my lessons learned into a reusable template for writing REST APIs, in the Go language.

It's mainly for myself. I don't want to keep reinventing the wheel and just want to get the foundation of my REST API 'ready to go' so I can focus on the business logic and integration with other systems and data stores.

Just to be clear: this is not a framework, library, package or anything like that. This tries to use a couple of very good Go packages and libraries that I like and cobbled them together.

The main ones are:

* [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) for routing
* [codegangsta/negroni](https://github.com/codegangsta/negroni) as a middleware handler
* [strechr/testify](https://github.com/stretchr/testify) for writing easier test assertions
* [unrolled/render](https://github.com/unrolled/render) for HTTP response rendering
* [palantir/stacktrace](https://github.com/palantir/stacktrace) to provide more context to error messages
* [unrolled/secure](https://github.com/unrolled/secure) to improve API security

### API Routes with Health check

```
Route{"Healthcheck", "GET", "/healthcheck", HealthcheckHandler},
//=== USERS ===
Route{"ListUsers", "GET", "/users", ListUsersHandler},
Route{"GetUser", "GET", "/users/{uid:[0-9]+}", GetUserHandler},
Route{"CreateUser", "POST", "/users", CreateUserHandler},
Route{"UpdateUser", "PUT", "/users/{uid:[0-9]+}", UpdateUserHandler},
Route{"DeleteUser", "DELETE", "/users/{uid:[0-9]+}", DeleteUserHandler},
```

## Testing

### Manually testing your API routes with curl commands

Let's start with some simple curl tests. Open your terminal and try the following curl commands.

Retrieve a list of users:

```
curl -X GET http://localhost:3009/users | python -mjson.tool
```

That should result in the following result:

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   216  100   216    0     0  14466      0 --:--:-- --:--:-- --:--:-- 15428
{
    "users": [
        {
            "dateOfBirth": "1992-01-01T00:00:00Z",
            "firstName": "Jane",
            "id": 1,
            "lastName": "Doe",
            "locationOfBirth": "Milton Keynes"
        },
        {
            "dateOfBirth": "1985-12-31T00:00:00Z",
            "firstName": "John",
            "id": 0,
            "lastName": "Doe",
            "locationOfBirth": "London"
        }
    ]
}
```

The `| python -mjson.tool` at the end is for pretty printing (formatting). It essentially tells to pipe the output of the curl command to the SJON formatting tool. If we only typed `curl -X GET http://localhost:3009/users` then we'd have something like this:

```
{"users":[{"id":0,"firstName":"John","lastName":"Doe","dateOfBirth":"1985-12-31T00:00:00Z","locationOfBirth":"London"},{"id":1,"firstName":"Jane","lastName":"Doe","dateOfBirth":"1992-01-01T00:00:00Z","locationOfBirth":"Milton Keynes"}]}
```

So not that easy to read as the earlier nicely formatted example.

Get a specific user:

```
curl -X GET http://localhost:3009/users/0 | python -mjson.tool
```

Results in:

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   104  100   104    0     0   6625      0 --:--:-- --:--:-- --:--:--  6933
{
    "dateOfBirth": "1992-01-01T00:00:00Z",
    "firstName": "Jane",
    "id": 1,
    "lastName": "Doe",
    "locationOfBirth": "Milton Keynes"
}
```

## Starting the app on a production server

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
