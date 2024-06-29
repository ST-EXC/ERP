// Package login is a generated protocol buffer package.
// It is generated from these files:
//
//	login.proto
//
// It has these top-level messages:
//
//	LoginRequest
//	LoginResponse
package login

// LoginRequest ...
type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

// LoginResponse ...
type LoginResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}
