// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package storage

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"

	printersinternal "github.com/onexstack/onex/internal/pkg/printers/internalversion"
	"github.com/onexstack/onex/internal/apiserver/registry/apps/miner"
	"github.com/onexstack/onex/pkg/apis/apps"
)

// MinerStorage includes storage for chains and all sub resources.
type MinerStorage struct {
	Miner  *REST
	Status *StatusREST
}

// NewStorage returns new instance of MinerStorage.
func NewStorage(optsGetter generic.RESTOptionsGetter) (MinerStorage, error) {
	minerRest, minerStatusRest, err := NewREST(optsGetter)
	if err != nil {
		return MinerStorage{}, err
	}

	return MinerStorage{
		Miner:  minerRest,
		Status: minerStatusRest,
	}, nil
}

// REST implements a RESTStorage for miners.
type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work against miners.
func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, *StatusREST, error) {
	store := &genericregistry.Store{
		NewFunc:       func() runtime.Object { return &apps.Miner{} },
		NewListFunc:   func() runtime.Object { return &apps.MinerList{} },
		PredicateFunc: miner.Matcher,
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*apps.Miner).Name, nil
		},
		DefaultQualifiedResource:  apps.Resource("miners"),
		SingularQualifiedResource: apps.Resource("miner"),

		CreateStrategy:      miner.Strategy,
		UpdateStrategy:      miner.Strategy,
		DeleteStrategy:      miner.Strategy,
		ResetFieldsStrategy: miner.Strategy,

		TableConvertor: printerstorage.TableConvertor{TableGenerator: printers.NewTableGenerator().With(printersinternal.AddHandlers)},
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: miner.GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, nil, err
	}

	// Subresources use the same store and creation strategy, which only
	// allows empty subs. Updates to an existing subresource are handled by
	// dedicated strategies.
	statusStore := *store
	statusStore.UpdateStrategy = miner.StatusStrategy
	statusStore.ResetFieldsStrategy = miner.StatusStrategy
	return &REST{store}, &StatusREST{store: &statusStore}, nil
}

// Implement ShortNamesProvider.
var _ rest.ShortNamesProvider = &REST{}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"mi"}
}

var _ rest.CategoriesProvider = &REST{}

// Categories implements the CategoriesProvider interface. Returns a list of categories a resource is part of.
func (r *REST) Categories() []string {
	return []string{"all"}
}

// StatusREST implements the REST endpoint for changing the status of a miner.
type StatusREST struct {
	store *genericregistry.Store
}

// New returns empty Chain object.
func (r *StatusREST) New() runtime.Object {
	return &apps.Miner{}
}

// Destroy cleans up resources on shutdown.
func (r *StatusREST) Destroy() {
	// Given that underlying store is shared with REST,
	// we don't destroy it here explicitly.
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return r.store.Get(ctx, name, options)
}

// Update alters the status subset of an object.
func (r *StatusREST) Update(
	ctx context.Context,
	name string,
	objInfo rest.UpdatedObjectInfo,
	createValidation rest.ValidateObjectFunc,
	updateValidation rest.ValidateObjectUpdateFunc,
	forceAllowCreate bool,
	options *metav1.UpdateOptions,
) (runtime.Object, bool, error) {
	// We are explicitly setting forceAllowCreate to false in the call to the underlying storage because
	// subresources should never allow create on update.
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, false, options)
}

// GetResetFields implements rest.ResetFieldsStrategy.
func (r *StatusREST) GetResetFields() map[fieldpath.APIVersion]*fieldpath.Set {
	return r.store.GetResetFields()
}

func (r *StatusREST) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return r.store.ConvertToTable(ctx, object, tableOptions)
}
