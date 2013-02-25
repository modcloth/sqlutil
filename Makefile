TARGETS := \
  null_types.go

all: build

build:
	go build -x $(TARGETS)

.PHONY: all build
