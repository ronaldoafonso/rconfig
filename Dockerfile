
FROM golang:alpine3.11

RUN apk add --update git postgresql-client

RUN addgroup rconfig && \
    adduser -h /home/rconfig -s /bin/ash -G rconfig -D rconfig

RUN mkdir -p /go/src/github.com && \
    chown -R rconfig:rconfig /go

USER rconfig:rconfig

ENV CGO_ENABLED 0

WORKDIR /go/src/github.com/ronaldoafonso/rconfig

COPY --chown=rconfig:rconfig . /go/src/github.com/ronaldoafonso/rconfig/

RUN go get -d -v ./...

RUN go install -v github.com/ronaldoafonso/rconfig

CMD ["/bin/sh", "-c", "while :; do sleep 10; done"]
