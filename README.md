# restapi-template-go
REST API server template written in Go.

This code is inspired by [How I write Go HTTP services after seven years][1].

It uses:
* [Gin-Gonic](https://github.com/gin-gonic/gin): HTTP web framework
* [GoConvey](https://github.com/smartystreets/goconvey): BDD style testing tool

## Requirements
* Go v1.11 or higher

## Build
```bash
# generate 'app-server' binary
$ make build

# build docker image
$ make build-image

# run docker image
$ make run-image
```

Server listens on `8080` port.

## Test
```bash
# Run test
$ make test

# Generate html coverage report
$ make cover-html

# Generate Cobertura coverage report
$ make cover-xml

# Generate JUnit report
$ make cover-xml-junit
```

* `reports/coverage.xml`: Cobertura coverage report
* `reports/junit.xml`: JUnit report

[1] https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831