// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// CSIVolumeSourceApplyConfiguration represents a declarative configuration of the CSIVolumeSource type for use
// with apply.
type CSIVolumeSourceApplyConfiguration struct {
	Driver               *string                                 `json:"driver,omitempty"`
	ReadOnly             *bool                                   `json:"readOnly,omitempty"`
	FSType               *string                                 `json:"fsType,omitempty"`
	VolumeAttributes     map[string]string                       `json:"volumeAttributes,omitempty"`
	NodePublishSecretRef *LocalObjectReferenceApplyConfiguration `json:"nodePublishSecretRef,omitempty"`
}

// CSIVolumeSourceApplyConfiguration constructs a declarative configuration of the CSIVolumeSource type for use with
// apply.
func CSIVolumeSource() *CSIVolumeSourceApplyConfiguration {
	return &CSIVolumeSourceApplyConfiguration{}
}

// WithDriver sets the Driver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Driver field is set to the value of the last call.
func (b *CSIVolumeSourceApplyConfiguration) WithDriver(value string) *CSIVolumeSourceApplyConfiguration {
	b.Driver = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *CSIVolumeSourceApplyConfiguration) WithReadOnly(value bool) *CSIVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *CSIVolumeSourceApplyConfiguration) WithFSType(value string) *CSIVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithVolumeAttributes puts the entries into the VolumeAttributes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the VolumeAttributes field,
// overwriting an existing map entries in VolumeAttributes field with the same key.
func (b *CSIVolumeSourceApplyConfiguration) WithVolumeAttributes(entries map[string]string) *CSIVolumeSourceApplyConfiguration {
	if b.VolumeAttributes == nil && len(entries) > 0 {
		b.VolumeAttributes = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.VolumeAttributes[k] = v
	}
	return b
}

// WithNodePublishSecretRef sets the NodePublishSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePublishSecretRef field is set to the value of the last call.
func (b *CSIVolumeSourceApplyConfiguration) WithNodePublishSecretRef(value *LocalObjectReferenceApplyConfiguration) *CSIVolumeSourceApplyConfiguration {
	b.NodePublishSecretRef = value
	return b
}
