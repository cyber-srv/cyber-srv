GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

.PHONY: proto
proto:
# gateway do not gen micro out
# gateway *.pb.go & *_gprc.go gen
	@protoc -I .. -I . -I ../../thirdparty/googleapis-master \
        --go_out=:. \
        --go_opt=Mservice-user/proto/service-user.proto=cyber-srv/gateway/proto/service-user \
		--go-grpc_out=:. \
        --go-grpc_opt=Mservice-user/proto/service-user.proto=cyber-srv/gateway/proto/service-user \
        --grpc-gateway_out=:. \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt generate_unbound_methods=true \
        --grpc-gateway_opt register_func_suffix=GW \
        --micro_out=:. \
        --micro_opt=Mservice-user/proto/service-user.proto=cyber-srv/gateway/proto/service-user \
        proto/gateway.proto
# service-user *.pb.go & *.pb.micro gen
	@protoc -I .. -I . \
        --go_out=:. \
        --micro_out=. \
        ../service-user/proto/service-user.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o gateway *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t gateway:latest .