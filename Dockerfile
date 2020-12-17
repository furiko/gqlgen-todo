FROM golang:1.15-alpine

COPY . /go/src/github.com/furiko/gqlgen-todo/
WORKDIR /go/src/github.com/furiko/gqlgen-todo/

CMD ["go", "run", "server.go"]