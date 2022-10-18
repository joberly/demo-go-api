# demo-go-api

A demonstration API service written in Go.

## Summary

The only endpoint is `/rand?min=<int64>&max=<int64>` which returns a random 64-bit integer between min and max so long as max > min >= 0. The `demo-cli` that can be build in the `cmd/demo-cli` directory is just an easy testing CLI which has named parameters for API query parameters. The user must run the API server `demo` in a separate terminal for `demo-cli` to function.

## Docker Image

`docker run --rm -it -p 8000:8000 ghcr.io/joberly/demo-go-api:latest`

### Sample Run

Server
```
> docker run --rm -it -p 8000:8000 ghcr.io/joberly/demo-go-api:latest
...
2022-10-18T05:12:52.640Z        INFO    demo/http.go:78 HTTP "Rand" mounted on GET /rand
2022-10-18T05:12:52.640Z        INFO    demo/http.go:87 HTTP server listening on ":8000"
2022-10-18T05:12:57.246Z        INFO    log/log.go:28   log     {"id": "pX3vgqbD", "req": "GET /rand?min=1&max=10", "from": "172.17.0.1"}
2022-10-18T05:12:57.246Z        INFO    demo-go-api/demo.go:31  demo.rand
2022-10-18T05:12:57.246Z        INFO    log/log.go:28   log     {"id": "pX3vgqbD", "status": 200, "bytes": 13, "time": "327.2µs"}
^C2022-10-18T05:13:45.172Z      INFO    demo/main.go:94 exiting (interrupt)
2022-10-18T05:13:45.172Z        INFO    demo/http.go:92 shutting down HTTP server at ":8000"
2022-10-18T05:13:45.172Z        INFO    demo/main.go:100        exited
```

Client
```
$ curl -vvv "http://localhost:8000/rand?min=1&max=10"
*   Trying ::1:8000...
* Connected to localhost (::1) port 8000 (#0)
> GET /rand?min=1&max=10 HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.70.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Tue, 18 Oct 2022 05:12:57 GMT
< Content-Length: 13
<
{"result":5}
* Connection #0 to host localhost left intact
```

## Local Setup

### Build

1. Install Golang verion >= 1.19.2.
2. Clone this repo.
3. `go build ./cmd/demo && go build ./cmd/demo-cli`

### Run

1. Run `./demo`
2. Open another terminal to the repo base directory to run `./demo-cli demo rand -h` for syntax.
   
### Example Run

Server:
```
> ./demo
2022-10-17T23:33:32.248-0500    INFO    demo/http.go:78 HTTP "Rand" mounted on GET /rand
2022-10-17T23:33:32.249-0500    INFO    demo/http.go:87 HTTP server listening on "localhost:8000"
2022-10-17T23:34:26.317-0500    INFO    log/log.go:28   log     {"id": "KXShm9aA", "req": "GET /rand?max=10&min=1", "from": "127.0.0.1"}
2022-10-17T23:34:26.318-0500    INFO    demo-go-api/demo.go:31  demo.rand
2022-10-17T23:34:26.318-0500    INFO    log/log.go:28   log     {"id": "KXShm9aA", "status": 200, "bytes": 13, "time": "1.0857ms"}
2022-10-17T23:43:43.241-0500    INFO    demo/main.go:94 exiting (interrupt)
2022-10-17T23:43:43.241-0500    INFO    demo/http.go:92 shutting down HTTP server at "localhost:8000"
2022-10-17T23:43:43.242-0500    INFO    demo/main.go:100        exited
```

Client
```
> ./demo-cli demo rand --min 1 --max 10
{
    "Result": 3
}
```
