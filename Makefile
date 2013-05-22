TARGETS := \
  github.com/modcloth/sqlutil

all: build

build:
	go build -x $(TARGETS)

test: build
	go test -x $(TARGETS)

.PHONY: all build
