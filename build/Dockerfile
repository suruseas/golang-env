
FROM golang:1.20.5-alpine

RUN apk update && apk add git \
                          make \
                          curl

RUN mkdir -p /go/src/github.com/suruseas/golang-env
WORKDIR /go/src/github.com/suruseas/golang-env

RUN go install golang.org/x/tools/cmd/goimports@latest

CMD ["ash"]
