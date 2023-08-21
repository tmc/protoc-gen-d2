# protoc-gen-d2

This project is a protoc plugin to render a d2 diagrams of protobuf packages.


Example:

![example](./testdata/acme/v1/test/test.erd.svg)

is rendered from

[./testdata/acme/v1/test/test.proto](./testdata/acme/v1/test/test.proto)
```proto
syntax = "proto3";

package acme.v1.test;

// Status enumeration.
enum Status {
  UNKNOWN = 0;
  ACTIVE = 1;
  INACTIVE = 2;
}

// Address nested message.
message Address {
  string street = 1;
  string city = 2;
  string state = 3;
  string country = 4;
}

// Sample message representing our test proto.
message TestMessage {
  string id = 1;
  string name = 2;
  repeated string items = 3;
  Status status = 4; // Enumeration field
  Address address = 5; // Nested message
  map<string, string> metadata = 6; // Map field
  oneof payload { // Oneof field
    string text_payload = 7;
    bytes binary_payload = 8;
  }
}

// Sample service for our test proto.
service TestService {
  // GetTestMessage retrieves a TestMessage
  rpc GetTestMessage (TestRequest) returns (TestMessage);
  // CreateTestMessage creates a new TestMessage.
  rpc CreateTestMessage (CreateRequest) returns (TestMessage);
}

// TestRequest is the request for GetTestMessage.
message TestRequest {
  string id = 1;
}

// CreateRequest is the request for CreateTestMessage.
message CreateRequest {
  TestMessage message = 1;
}
```