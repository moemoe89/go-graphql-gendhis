FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-graphql-gendhis

WORKDIR /go/src/github.com/moemoe89/go-graphql-gendhis

COPY . /go/src/github.com/moemoe89/go-graphql-gendhis

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/goose -env=docker up && /go/bin/go-graphql-gendhis
