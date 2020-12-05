generate:
	protoc -I . \
	 --go_out=:. \
	 --go_opt=paths=source_relative \
	 --plugin=${GOPATH}/bin/version/v1.25.0/protoc-gen-go \
	 redact/redact.proto

fmt:
	GO111MODULE=on go fmt .

lint: fmt
	GO111MODULE=on go vet --vettool=${GOPATH}/bin/shadow .
	staticcheck .

clean: lint
	GO111MODULE=on go mod tidy
	rm -rf bin/*

build: clean
	GO111MODULE=on go build -o bin/protoc-gen-redact .

examples: build
	protoc -I . \
	 --go_out=:. \
	 --go_opt=paths=source_relative \
	 --plugin=${GOPATH}/bin/version/v1.25.0/protoc-gen-go \
	 --go-grpc_out=:. \
	 --go-grpc_opt=paths=source_relative \
	 --redact_out=:. \
	 --redact_opt=paths=source_relative \
	 --plugin=bin/protoc-gen-redact \
	 examples/user/pb/user.proto \
	 examples/tests/message.proto

.PHONY: generate fmt lint clean build examples
