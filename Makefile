.EXPORT_ALL_VARIABLES:
GOBIN = $(shell pwd)/bin
GO111MODULE = on
SHELL=/bin/bash

.PHONY: deps
deps:
	@go mod vendor
	@go mod tidy
	@./proto-third-party.sh

.PHONY: tools
tools:
	@cd tools ; ./install.sh

.PHONY: clean
clean:
	@rm -fv ./bin/*

.PHONY: generate
generate: deps tools
	@./generate.sh

#
# Protobuf
#

# PROTOC is required to generate code from proto files. If make target depends
# on it and it is not installed, make will be terminated.
PROTOC := $(shell which protoc)
ifeq ($(PROTOC),)
	PROTOC = must-install-protoc
endif

$(PROTOC):
	@echo -e "\nERROR: protobuf is not installed. Use your package manager to install or compile from source. See https://github.com/protocolbuffers/protobuf for details.\n"
	@exit 1
