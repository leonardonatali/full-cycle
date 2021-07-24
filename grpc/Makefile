# Export all variables in env file
include .env
export $(shell sed 's/=.*//' .env)

default: run

protoc:
	@protoc \
		--proto_path=proto \
		proto/*.proto \
		--go_out=pkg/protobuf \
		--go-grpc_out=pkg/protobuf
run:
	@go run cmd/main.go 