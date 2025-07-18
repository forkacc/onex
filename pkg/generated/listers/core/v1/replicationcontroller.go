// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ReplicationControllerLister helps list ReplicationControllers.
// All objects returned here must be treated as read-only.
type ReplicationControllerLister interface {
	// List lists all ReplicationControllers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*corev1.ReplicationController, err error)
	// ReplicationControllers returns an object that can list and get ReplicationControllers.
	ReplicationControllers(namespace string) ReplicationControllerNamespaceLister
	ReplicationControllerListerExpansion
}

// replicationControllerLister implements the ReplicationControllerLister interface.
type replicationControllerLister struct {
	listers.ResourceIndexer[*corev1.ReplicationController]
}

// NewReplicationControllerLister returns a new ReplicationControllerLister.
func NewReplicationControllerLister(indexer cache.Indexer) ReplicationControllerLister {
	return &replicationControllerLister{listers.New[*corev1.ReplicationController](indexer, corev1.Resource("replicationcontroller"))}
}

// ReplicationControllers returns an object that can list and get ReplicationControllers.
func (s *replicationControllerLister) ReplicationControllers(namespace string) ReplicationControllerNamespaceLister {
	return replicationControllerNamespaceLister{listers.NewNamespaced[*corev1.ReplicationController](s.ResourceIndexer, namespace)}
}

// ReplicationControllerNamespaceLister helps list and get ReplicationControllers.
// All objects returned here must be treated as read-only.
type ReplicationControllerNamespaceLister interface {
	// List lists all ReplicationControllers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*corev1.ReplicationController, err error)
	// Get retrieves the ReplicationController from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*corev1.ReplicationController, error)
	ReplicationControllerNamespaceListerExpansion
}

// replicationControllerNamespaceLister implements the ReplicationControllerNamespaceLister
// interface.
type replicationControllerNamespaceLister struct {
	listers.ResourceIndexer[*corev1.ReplicationController]
}
