BINARY := etcdjsonparser
ETCD_JSON_PARSER_VERSION := 1.0.0
REV_VAR := etcd-json-parser/common/version.FullVersion
GOBUILD_VERSION_ARGS := -ldflags "-X $(REV_VAR)=$(ETCD_JSON_PARSER_VERSION)"
ROOT := $(PWD)
GOPATH ?= $(ROOT)/../..
SHELL := /bin/bash

.PHONY: all
all: build

.PHONY: build
build:
	GOPATH=$(GOPATH) GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags netgo -installsuffix netgo -o $(GOPATH)/bin/$(BINARY) $(GOBUILD_VERSION_ARGS) .

.PHONY: clean
clean:
	rm tst

