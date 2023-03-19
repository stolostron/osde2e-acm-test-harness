DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
OUT_FILE := "$(DIR)osde2e-acm-test-harness"

REPOSITORY := ${REPOSITORY}
IMAGE := osde2e-acm-test-harness
VERSION := latest

# to ignore vendor directory
GOFLAGS=-mod=mod
build:
	CGO_ENABLED=0 go test -v -c

lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.46.2
	(cd "$(DIR)"; golangci-lint run -c .ci-operator.yaml ./...)

build-e2e:
	docker build . -t ${REPOSITORY}/${IMAGE}:${VERSION}
	docker push ${REPOSITORY}/${IMAGE}:${VERSION}

test-local: build-e2e
	@echo "Starting e2e test local"
	./hack/run-acm-addon-harness.sh

