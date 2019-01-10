BASE_DIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
REPORT_DIR=reports
TEST_DIR=./app/...

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOCOVER=$(GOCMD) tool cover
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

ENV_GOMOD_ON=GO111MODULE=on
ENV_TEST=
# uncomment to use 'vendor' directory
#VENDOR_OPT=-mod=vendor
GOBUILD_OPT=$(VENDOR_OPT) -v
GOTEST_OPT=$(VENDOR_OPT) -v

BINARY=app-server
COVER_HTML=$(REPORT_DIR)/coverage.html
COVER_TXT=$(REPORT_DIR)/coverage.txt
COVER_XML=$(REPORT_DIR)/coverage.xml
JUNIT_REPORT=$(REPORT_DIR)/junit.xml

VERSION=latest
DOCKER_TAG=restapi-template-go:$(VERSION)


all: test build

report-init:
	@mkdir -p $(REPORT_DIR)
install-cover-tool:
	@$(GOCMD) get github.com/t-yuki/gocover-cobertura
install-junit-tool:
	@$(GOCMD) get -u github.com/jstemmer/go-junit-report

# Build
build:
	@$(ENV_GOMOD_ON) $(GOBUILD) $(GOBUILD_OPT) -o $(BINARY)

build-image:
	@docker build -t $(DOCKER_TAG) .

# Test
test:
	@$(ENV_GOMOD_ON) $(ENV_TEST) $(GOTEST) $(GOTEST_OPT) -count=1 $(TEST_DIR)

cover: report-init
	@$(ENV_GOMOD_ON) $(GOTEST) $(GOTEST_OPT) -covermode=count -coverprofile=$(COVER_TXT) -coverpkg=$(TEST_DIR) $(TEST_DIR)
cover-html: cover
	@$(GOCOVER) -html=$(COVER_TXT) -o $(COVER_HTML)
cover-cout: cover
	@$(GOCOVER) -func=$(COVER_TXT)
cover-xml: install-cover-tool cover
	@gocover-cobertura < $(COVER_TXT) > $(COVER_XML)

cover-xml-junit: report-init install-junit-tool
	@make cover-xml | go-junit-report > $(JUNIT_REPORT)

# Clean
clean:
	@$(GOCLEAN)
	@rm -f $(BINARY) $(COVER_TXT) $(COVER_HTML) $(COVER_XML) $(JUNIT_REPORT)

# Run
run: build
	@./$(BINARY)

run-image:
	@docker run --rm -p8080:8080 $(DOCKER_TAG)

# Install dependencies to vendor/
vendor:
	@$(GOMOD) vendor
vendor-update:
	@$(GOGET) -u
