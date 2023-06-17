# Log Service

A super simple logger that accepts a string and appends to an `app.log` text file.

## Get started

Build the log service using the [service handler](../service/)

```shell
go build ./cmd/logservice
```

Then start the logservice and create logs via an HTTP `POST` request.

```shell
./logservice

curl --request POST \
    --url http://localhost:4000/log \
    --data "<Message>"
```

