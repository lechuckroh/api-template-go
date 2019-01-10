[![GitHub license](https://img.shields.io/github/license/r0oth3x49/udemy-dl.svg?style=flat-square)](https://github.com/lechuckroh/api-template-go/blob/master/LICENSE)
[![Build Status](https://travis-ci.com/lechuckroh/api-template-go.svg?branch=master)](https://travis-ci.com/lechuckroh/api-template-go)

# api-template-go

HTTP API server template written in Go.

This code is inspired by [How I write Go HTTP services after seven years](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831).

It uses:
* [Gin-Gonic](https://github.com/gin-gonic/gin): HTTP web framework
* [GoConvey](https://github.com/smartystreets/goconvey): BDD style testing tool

## Requirements
* Go 1.11 or higher

## Build
```bash
# generate 'app-server' binary
$ make build

# build docker image
$ make build-image

# build docker image using docker
$ make docker-build-image

# run docker image
$ make run-image
```

Server listens on `8080` port.

## Test
```bash
# Run test
$ make test

# Generate coverage report
$ make cover

# Generate coverage report in html format
$ make cover-html
```

* `reports/coverage.xml`: Cobertura coverage report
* `reports/junit.xml`: JUnit report
