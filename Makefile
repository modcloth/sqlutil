TARGETS := \
  github.com/modcloth/sqlutil

all: build test

build:
	go build -x $(TARGETS)

deps:
	go get -x github.com/golang/lint/golint

test: build fmtpolice
	go test -x $(TARGETS)

fmtpolice:
	set -e ; for f in $(shell git ls-files '*.go'); do gofmt $$f | diff -u $$f - ; done
	set -e ; for f in $(shell git ls-files '*.go'); do $${GOPATH%%:*}/bin/golint $$f ; done

.PHONY: all build test fmtpolice
