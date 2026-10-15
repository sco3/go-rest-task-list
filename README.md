# Simple Go HTTP Task API

A lightweight REST API built with Go's standard `net/http` library 
for managing in-memory tasks with thread-safe synchronization.

Run the server with `go run main.go`, then interact with it using curl (list tasks):

```sh
curl localhost:8080/tasks
```

Create tasks:

```sh
curl -X POST localhost:8080/tasks \
   -H "Content-Type: application/json" \
   -d '{"title":"New Task"}'
```