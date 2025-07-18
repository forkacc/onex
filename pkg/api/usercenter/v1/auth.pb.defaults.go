// This file defines the Protobuf messages for managing Users.
//

// Code generated by protoc-gen-defaults. DO NOT EDIT.

package v1

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	_ *timestamppb.Timestamp
	_ *durationpb.Duration
	_ *wrapperspb.BoolValue
)

func (x *AuthenticateRequest) Default() {
}

func (x *AuthenticateResponse) Default() {
}

func (x *AuthorizeRequest) Default() {
}

func (x *AuthorizeResponse) Default() {
}

func (x *AuthRequest) Default() {
}

func (x *AuthResponse) Default() {
}
