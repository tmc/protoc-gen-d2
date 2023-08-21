# protoc-gen-d2

This project is a protoc plugin to render a d2 diagrams of protobuf packages.


Example:

![example](./testdata/acme/test/v1/test.erd.svg)

is rendered from

[./testdata/acme/test/v1/test.proto](./testdata/acme/test/v1/test.proto)
```proto

syntax = "proto3";

package acme.test.v1;

// Sample service for our test proto.
service TestService {
  // GetTestMessage retrieves a TestMessage
  rpc GetTestMessage (GetTestMessageRequest) returns (GetTestMessageResponse);
  // CreateTestMessage creates a new TestMessage.
  rpc CreateTestMessage (CreateTestMessageRequest) returns (CreateTestMessageResponse);
}

// Status enumeration.
enum Status {
  STATUS_UNSPECIFIED = 0;
  STATUS_ACTIVE = 1;
  STATUS_INACTIVE = 2;
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


// GetTestMessageRequest is the request for GetTestMessage.
message GetTestMessageRequest {
  string id = 1;
}

// GetTestMessageResponse is the response for GetTestMessage.
message GetTestMessageResponse {
  TestMessage message = 1;
}

// CreateTestMessageRequest is the request for CreateTestMessage.
message CreateTestMessageRequest {
  TestMessage message = 1;
}

// CreateTestMessageResponse is the response for CreateTestMessage.
message CreateTestMessageResponse {
  TestMessage message = 1;
}
```
