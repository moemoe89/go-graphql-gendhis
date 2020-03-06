FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-graphql-golang

WORKDIR /go/src/github.com/moemoe89/practicing-graphql-golang

COPY . /go/src/github.com/moemoe89/practicing-graphql-golang

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/goose -env=docker up && /go/bin/practicing-graphql-golang
