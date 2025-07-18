// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// DownwardAPIVolumeSourceApplyConfiguration represents a declarative configuration of the DownwardAPIVolumeSource type for use
// with apply.
type DownwardAPIVolumeSourceApplyConfiguration struct {
	Items       []DownwardAPIVolumeFileApplyConfiguration `json:"items,omitempty"`
	DefaultMode *int32                                    `json:"defaultMode,omitempty"`
}

// DownwardAPIVolumeSourceApplyConfiguration constructs a declarative configuration of the DownwardAPIVolumeSource type for use with
// apply.
func DownwardAPIVolumeSource() *DownwardAPIVolumeSourceApplyConfiguration {
	return &DownwardAPIVolumeSourceApplyConfiguration{}
}

// WithItems adds the given value to the Items field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Items field.
func (b *DownwardAPIVolumeSourceApplyConfiguration) WithItems(values ...*DownwardAPIVolumeFileApplyConfiguration) *DownwardAPIVolumeSourceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithItems")
		}
		b.Items = append(b.Items, *values[i])
	}
	return b
}

// WithDefaultMode sets the DefaultMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultMode field is set to the value of the last call.
func (b *DownwardAPIVolumeSourceApplyConfiguration) WithDefaultMode(value int32) *DownwardAPIVolumeSourceApplyConfiguration {
	b.DefaultMode = &value
	return b
}
