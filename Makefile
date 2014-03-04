TARGETS := \
  github.com/modcloth/sqlutil

all: build test

build:
	go build -x $(TARGETS)

test:
	go test -x $(TARGETS)

.PHONY: all build test
