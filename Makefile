generate:
	protoc -I . \
	 --go_out=:. \
	 --go_opt=paths=source_relative \
	 --plugin=${GOPATH}/bin/version/v1.25.0/protoc-gen-go \
	 redact/redact.proto

.PHONY: generate
