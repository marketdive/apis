#!/usr/bin/env bash
set -euo pipefail

root_dir="$(pwd)"

# Run this script with FORCE like `FORCE=1 ./hack/proto-third-party.sh` to re-download everything.
FORCE="${FORCE:-}"

# Google common protobuf definitions

if [[ -n "${FORCE}" || ! -r "third_party/google/api/annotations.proto" ]]; then
  repo="https://github.com/googleapis/api-common-protos.git"

  echo "Cloning ${repo} ..."
  rm -rf /tmp/api-common-protos/
  cd /tmp
  git clone -q ${repo}
  cd api-common-protos
  git checkout -q 1.50.0

  echo "Copying Google proto files to third_party/ ..."
  cd /tmp/api-common-protos
  mkdir -p "${root_dir}/third_party/google/api/"
  mkdir -p "${root_dir}/third_party/google/type/"
  cp google/api/{annotations,field_behavior,http,httpbody}.proto "${root_dir}/third_party/google/api/"
  cp google/type/date.proto "${root_dir}/third_party/google/type/"
fi

# grpc-gateway protobuf definitions

if [[ -n "${FORCE}" || ! -r "third_party/protoc-gen-swagger/options/annotations.proto" ]]; then
  echo "Copying grpc-gateway & swagger proto files to third_party/ ..."
  cd "${root_dir}/vendor/github.com/grpc-ecosystem/grpc-gateway"
  mkdir -p "${root_dir}/third_party/protoc-gen-swagger/options/"
  cp protoc-gen-swagger/options/{annotations,openapiv2}.proto "${root_dir}/third_party/protoc-gen-swagger/options/"
fi
