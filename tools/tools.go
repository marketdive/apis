// +build tools

package main

// This file defines tool dependencies for the module.
// See: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

import (
	// Base plugins for Protobuf & gRPC
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	// gRPC-gateway & swagger plugins
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	// Other plugins
	_ "github.com/TheThingsIndustries/protoc-gen-fieldmask"
	// Development tools
	_ "github.com/uber/prototool/cmd/prototool"
)
