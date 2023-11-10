database ?= "in-memory"

run:
	go run ./cmd/main.go -database=$(database)

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

test:
	go test ./...

.PHONY: run proto test