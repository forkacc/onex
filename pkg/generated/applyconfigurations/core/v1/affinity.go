// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// AffinityApplyConfiguration represents a declarative configuration of the Affinity type for use
// with apply.
type AffinityApplyConfiguration struct {
	NodeAffinity    *NodeAffinityApplyConfiguration    `json:"nodeAffinity,omitempty"`
	PodAffinity     *PodAffinityApplyConfiguration     `json:"podAffinity,omitempty"`
	PodAntiAffinity *PodAntiAffinityApplyConfiguration `json:"podAntiAffinity,omitempty"`
}

// AffinityApplyConfiguration constructs a declarative configuration of the Affinity type for use with
// apply.
func Affinity() *AffinityApplyConfiguration {
	return &AffinityApplyConfiguration{}
}

// WithNodeAffinity sets the NodeAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeAffinity field is set to the value of the last call.
func (b *AffinityApplyConfiguration) WithNodeAffinity(value *NodeAffinityApplyConfiguration) *AffinityApplyConfiguration {
	b.NodeAffinity = value
	return b
}

// WithPodAffinity sets the PodAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodAffinity field is set to the value of the last call.
func (b *AffinityApplyConfiguration) WithPodAffinity(value *PodAffinityApplyConfiguration) *AffinityApplyConfiguration {
	b.PodAffinity = value
	return b
}

// WithPodAntiAffinity sets the PodAntiAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodAntiAffinity field is set to the value of the last call.
func (b *AffinityApplyConfiguration) WithPodAntiAffinity(value *PodAntiAffinityApplyConfiguration) *AffinityApplyConfiguration {
	b.PodAntiAffinity = value
	return b
}
