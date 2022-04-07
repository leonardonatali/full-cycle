# Export all variables in env file
include .env
export $(shell sed 's/=.*//' .env)

default: run

install:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

protoc:
	@protoc \
		--proto_path=proto \
		proto/*.proto \
		--go_out=pkg/protobuf \
		--go-grpc_out=pkg/protobuf
run:
	@go run cmd/main.go 