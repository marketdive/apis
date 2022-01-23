module github.com/marketdive/apis/tools

go 1.16

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.5.1-gogo.1

replace github.com/envoyproxy/protoc-gen-validate => github.com/TheThingsIndustries/protoc-gen-validate v0.4.0-fieldmask.1

// See: https://github.com/uber/prototool/issues/559
replace github.com/fullstorydev/grpcurl => github.com/fullstorydev/grpcurl v1.8.0

require (
	github.com/TheThingsIndustries/protoc-gen-fieldmask v0.4.4
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/uber/prototool v1.10.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.23.0
)
