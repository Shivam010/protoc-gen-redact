protoc-gen-redact (PGR)
=======================
[![Go Reference](https://pkg.go.dev/badge/Shivam010/protoc-gen-redact.svg)](https://pkg.go.dev/Shivam010/protoc-gen-redact)
[![License](https://img.shields.io/badge/license-apache2-mildgreen.svg)](./LICENSE)
[![GitHub release](https://img.shields.io/github/release/Shivam010/protoc-gen-redact.svg)](https://github.com/Shivam010/protoc-gen-redact/releases)

_protoc-gen-redact (PGR)_ is a protoc plugin to redact field values in GRPC client calls from the server. This plugin
adds support to protoc-generated code to redact certain fields in the GRPC calls.

Developers only need import the PGR extension and annotate the messages or fields in their proto files to redact:

```protobuf
syntax = "proto3";

package user;

import "redact/redact.proto";
import "google/protobuf/empty.proto";

message User {
    // User credentials
    string username = 1;
    string password = 2 [(redact.redact) = true];

    // User information
    string email = 3 [(redact.custom).string = "r*d@ct*d"];
    string name = 4;
    Location home = 5;
    message Location {
        double lat = 1;
        double lng = 2;
    }
}

service Chat {
    rpc GetUser(GetUserRequest) returns (User);
    rpc GetUserInternal(GetUserRequest) returns (User) {
        option (redact.method_skip) = true;
    }
    rpc ListUsers (google.protobuf.Empty) returns (ListUsersResponse) {
        option (redact.internal_method) = true;
    }
}

message GetUserRequest {
    string username = 1;
}

message ListUsersResponse {
    repeated User users = 1;
}

```

Request for Contribution
------------------------
Contributors are more than welcome and much appreciated. Please feel free to open a PR to improve anything you don't
like, or would like to add.

Please make your changes in a specific branch and create a pull request into master! If you can, please make sure all
the changes work properly and does not affect the existing functioning.

No PR is too small! Even the smallest effort is countable.

License
-------
This project is licensed under the [Apache License 2.0](./LICENSE)
