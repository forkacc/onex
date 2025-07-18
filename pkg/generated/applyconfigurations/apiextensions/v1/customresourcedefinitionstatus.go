// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// CustomResourceDefinitionStatusApplyConfiguration represents a declarative configuration of the CustomResourceDefinitionStatus type for use
// with apply.
type CustomResourceDefinitionStatusApplyConfiguration struct {
	Conditions     []CustomResourceDefinitionConditionApplyConfiguration `json:"conditions,omitempty"`
	AcceptedNames  *CustomResourceDefinitionNamesApplyConfiguration      `json:"acceptedNames,omitempty"`
	StoredVersions []string                                              `json:"storedVersions,omitempty"`
}

// CustomResourceDefinitionStatusApplyConfiguration constructs a declarative configuration of the CustomResourceDefinitionStatus type for use with
// apply.
func CustomResourceDefinitionStatus() *CustomResourceDefinitionStatusApplyConfiguration {
	return &CustomResourceDefinitionStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *CustomResourceDefinitionStatusApplyConfiguration) WithConditions(values ...*CustomResourceDefinitionConditionApplyConfiguration) *CustomResourceDefinitionStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithAcceptedNames sets the AcceptedNames field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AcceptedNames field is set to the value of the last call.
func (b *CustomResourceDefinitionStatusApplyConfiguration) WithAcceptedNames(value *CustomResourceDefinitionNamesApplyConfiguration) *CustomResourceDefinitionStatusApplyConfiguration {
	b.AcceptedNames = value
	return b
}

// WithStoredVersions adds the given value to the StoredVersions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the StoredVersions field.
func (b *CustomResourceDefinitionStatusApplyConfiguration) WithStoredVersions(values ...string) *CustomResourceDefinitionStatusApplyConfiguration {
	for i := range values {
		b.StoredVersions = append(b.StoredVersions, values[i])
	}
	return b
}
