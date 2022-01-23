#!/usr/bin/env bash
set -euo pipefail

# inner/wildberries/v1
protoc -I . \
  -I third_party/ \
  --plugin=./bin/protoc-gen-go \
  --go_out=. \
  --go_opt=paths=source_relative \
  --plugin=./bin/protoc-gen-go-grpc \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  inner/wildberries/v1/*.proto
