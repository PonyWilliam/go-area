
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v2/cmd/protoc-gen-micro
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/area.proto
	ls proto/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
.PHONY: build
build:
	go build -o product *.go

.PHONY: test
test:
	go test -v ./... -cover

docker:
	docker build -t ponywilliam/go-area .
	docker tag ponywilliam/go-area ponywilliam/go-area
	docker push ponywilliam/go-area