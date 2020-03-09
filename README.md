[![Build Status](https://travis-ci.org/moemoe89/practicing-graphql-golang.svg?branch=master)](https://travis-ci.org/moemoe89/practicing-graphql-golang)
[![codecov](https://codecov.io/gh/moemoe89/practicing-graphql-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/moemoe89/practicing-graphql-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/practicing-graphql-golang)](https://goreportcard.com/report/github.com/moemoe89/practicing-graphql-golang)

# practicing-graphql-golang #

Simple Go Clean Arch Using Golang (Gin Gonic Framework) as Programming Language, PostgreSQL as Database

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ practicing-graphql-golang/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup PostgreSQL <https://www.postgresql.org/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> practicing-graphql-golang
```

## Running Migration
* Copy `config-sample.json` to `config.json` and changes the value based on your configurations.
* Create PostgreSQL database for example named `simple_api` and do migration with `Goose` <https://bitbucket.org/liamstask/goose/>
* Change database configuration on dbconf.yml like `dialect` and `dsn` for each environtment
* Do the following command
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-graphql-golang
$ goose -env=development up
```

## Documetation with Swagger
For open swagger access via browser
```
{{url}}/swagger/index.html
```
For updating swagger
```
$ swag init
```

## Running Application with Makefile
Make config file for local :
```
$ cp config-sample.json config-local.json
```
Build
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-graphql-golang
$ make build
```
Run
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-graphql-golang
$ make run
```
Stop
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-graphql-golang
$ make stop
```
Make config file for docker :
```
$ cp config-sample.json config-docker.json
```
Docker Build
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-graphql-golang
$ make docker-build
```
Docker Up
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-graphql-golang
$ make docker-up
```
Docker Down
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-graphql-golang
$ make docker-down
```

## How to Run with Docker
Make config file for docker :
```
$ cp config-sample.json config.json
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```

## How to Run Unit Test
Run
```
$ go test ./...
```
Run with cover
```
$ go test ./... -cover
```
Run with HTML output
```
$ go test ./... -coverprofile=c.out && go tool cover -html=c.out
```

## Example Request
Navigate your browser to this url for running the GraphQL console
```
/api/v1/graphql/user
```
### Create
```
POST /api/v1/graphql/user
Content-Type: application/json
{
	"query": "mutation{Create(name:\"momo\",phone:\"0856\",email:\"m@m.com\",address:\"Indonesia\"){id,name,phone,email,address}}"
}
```
### List
```
POST /api/v1/graphql/user
Content-Type: application/json
{
	"query": "{List(per_page:\"10\",page:\"1\",order_by:\"\",name:\"\",phone:\"\",email:\"\",created_at_start:\"\",created_at_end:\"\",select_field:\"\"){list{id,name,phone,email,address}}}"
}
```
### Detail
```
POST /api/v1/graphql/user
Content-Type: application/json
{
	"query": "{Detail(id:\"bph2mlript32plmed820\"){id,name,phone,email,address}}"
}
```
### Update
```
POST /api/v1/graphql/user
Content-Type: application/json
{
	"query": "mutation{Update(id:\"bpielbbipt341rif5i20\",name:\"momo update\",phone:\"0856\",email:\"m@m.com\",address:\"Indonesia\"){id,name,phone,email,address}}"
}
```
### Delete
```
POST /api/v1/graphql/user
Content-Type: application/json
{
	"query": "mutation{Delete(id:\"bpielbbipt341rif5i20\")}"
}
```

## Reference

Thanks to this medium [link](https://medium.com/easyread/graphql-delivery-on-golangs-clean-architecture-5c995a17b3a8) for sharing the great article

## License

MIT