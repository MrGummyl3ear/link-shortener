database ?= "in-memory"

run:
	go run ./cmd/main.go -database=$(database)

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

df:
	sudo docker build --build-arg DATABASE=$(database) --tag link-shortener .

service_up:
	sudo docker-compose -f docker-compose.yml up -d --remove-orphans

test:
	go test ./...

.PHONY: run proto df service_up test