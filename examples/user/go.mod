module github.com/Shivam010/protoc-gen-redact/examples/user

go 1.14

require (
	github.com/Shivam010/protoc-gen-redact v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/Shivam010/protoc-gen-redact => ../..
