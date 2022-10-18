# TODO cross-compile executable with native arch builder
ARG ARCH=
FROM ${ARCH}golang:1.19.2 AS builder

COPY ./ /go/src/github.com/joberly/demo-go-api

RUN set -e ;\
    cd /go/src/github.com/joberly/demo-go-api ;\
    go test -timeout 30s ;\
    cd /go/src/github.com/joberly/demo-go-api/cmd/demo ;\
    go build .

FROM ${ARCH}debian:stable-slim

COPY --from=builder /go/src/github.com/joberly/demo-go-api/cmd/demo/demo /

EXPOSE 8000/tcp
ENTRYPOINT "/demo"
