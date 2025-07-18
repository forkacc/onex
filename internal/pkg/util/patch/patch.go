// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package patch

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	"github.com/onexstack/onex/internal/pkg/util/conditions"
	coreutil "github.com/onexstack/onex/internal/pkg/util/core"
	"github.com/onexstack/onex/pkg/apis/apps/v1beta1"
)

// Helper is a utility for ensuring the proper patching of objects.
type Helper struct {
	client       client.Client
	gvk          schema.GroupVersionKind
	beforeObject client.Object
	before       *unstructured.Unstructured
	after        *unstructured.Unstructured
	changes      map[string]bool

	isConditionsSetter bool
}

// NewHelper returns an initialized Helper. Use NewHelper before changing obj.
// After changing obj use Helper.Patch to persist your changes.
func NewHelper(obj client.Object, crClient client.Client) (*Helper, error) {
	// Return early if the object is nil.
	if coreutil.IsNil(obj) {
		return nil, errors.New("helper could not be created: object is nil")
	}

	// Get the GroupVersionKind of the object,
	// used to validate against later on.
	gvk, err := apiutil.GVKForObject(obj, crClient.Scheme())
	if err != nil {
		return nil, err
	}

	// Convert the object to unstructured to compare against our before copy.
	unstructuredObj, err := toUnstructured(obj)
	if err != nil {
		return nil, err
	}

	// Check if the object satisfies the Cluster API conditions contract.
	_, canInterfaceConditions := obj.(conditions.Setter)

	return &Helper{
		client:             crClient,
		gvk:                gvk,
		before:             unstructuredObj,
		beforeObject:       obj.DeepCopyObject().(client.Object),
		isConditionsSetter: canInterfaceConditions,
	}, nil
}

// Patch will attempt to patch the given object, including its status.
func (h *Helper) Patch(ctx context.Context, obj client.Object, opts ...Option) error {
	// Return early if the object is nil.
	if coreutil.IsNil(obj) {
		return errors.New("Patch could not be completed: object is nil")
	}

	// Get the GroupVersionKind of the object that we want to patch.
	gvk, err := apiutil.GVKForObject(obj, h.client.Scheme())
	if err != nil {
		return err
	}
	if gvk != h.gvk {
		return errors.Errorf("unmatched GroupVersionKind, expected %q got %q", h.gvk, gvk)
	}

	// Calculate the options.
	options := &HelperOptions{}
	for _, opt := range opts {
		opt.ApplyToHelper(options)
	}

	// Convert the object to unstructured to compare against our before copy.
	h.after, err = toUnstructured(obj)
	if err != nil {
		return err
	}

	// Determine if the object has status.
	if unstructuredHasStatus(h.after) {
		if options.IncludeStatusObservedGeneration {
			// Set status.observedGeneration if we're asked to do so.
			if err := unstructured.SetNestedField(h.after.Object, h.after.GetGeneration(), "status", "observedGeneration"); err != nil {
				return err
			}

			// Restore the changes back to the original object.
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(h.after.Object, obj); err != nil {
				return err
			}
		}
	}

	// Calculate and store the top-level field changes (e.g. "metadata", "spec", "status") we have before/after.
	h.changes, err = h.calculateChanges(obj)
	if err != nil {
		return err
	}

	// If it is a deletion reconcile, only patch once, to avoid `not found` error when patch status.
	// In a deletion reconcile, we don't need patch status.
	if !obj.GetDeletionTimestamp().IsZero() {
		return kerrors.NewAggregate([]error{h.patch(ctx, obj)})
	}

	// Issue patches and return errors in an aggregate.
	return kerrors.NewAggregate([]error{
		// Patch the conditions first.
		//
		// Given that we pass in metadata.resourceVersion to perform a 3-way-merge conflict resolution,
		// patching conditions first avoids an extra loop if spec or status patch succeeds first
		// given that causes the resourceVersion to mutate.
		h.patchStatusConditions(ctx, obj, options.ForceOverwriteConditions, options.OwnedConditions),

		// Then proceed to patch the rest of the object.
		h.patch(ctx, obj),
		h.patchStatus(ctx, obj),
	})
}

// patch issues a patch for metadata and spec.
func (h *Helper) patch(ctx context.Context, obj client.Object) error {
	if !h.shouldPatch("metadata") && !h.shouldPatch("spec") {
		return nil
	}
	beforeObject, afterObject, err := h.calculatePatch(obj, specPatch)
	if err != nil {
		return err
	}
	return h.client.Patch(ctx, afterObject, client.MergeFrom(beforeObject))
}

// patchStatus issues a patch if the status has changed.
func (h *Helper) patchStatus(ctx context.Context, obj client.Object) error {
	if !h.shouldPatch("status") {
		return nil
	}
	beforeObject, afterObject, err := h.calculatePatch(obj, statusPatch)
	if err != nil {
		return err
	}
	return h.client.Status().Patch(ctx, afterObject, client.MergeFrom(beforeObject))
}

// patchStatusConditions issues a patch if there are any changes to the conditions slice under
// the status subresource. This is a special case and it's handled separately given that
// we allow different controllers to act on conditions of the same object.
//
// This method has an internal backoff loop. When a conflict is detected, the method
// asks the Client for the a new version of the object we're trying to patch.
//
// Condition changes are then applied to the latest version of the object, and if there are
// no unresolvable conflicts, the patch is sent again.
func (h *Helper) patchStatusConditions(ctx context.Context, obj client.Object, forceOverwrite bool, ownedConditions []v1beta1.ConditionType) error {
	// Nothing to do if the object isn't a condition patcher.
	if !h.isConditionsSetter {
		return nil
	}

	// Make sure our before/after objects satisfy the proper interface before continuing.
	//
	// NOTE: The checks and error below are done so that we don't panic if any of the objects don't satisfy the
	// interface any longer, although this shouldn't happen because we already check when creating the patcher.
	before, ok := h.beforeObject.(conditions.Getter)
	if !ok {
		return errors.Errorf("object %s doesn't satisfy conditions.Getter, cannot patch", before.GetObjectKind())
	}
	after, ok := obj.(conditions.Getter)
	if !ok {
		return errors.Errorf("object %s doesn't satisfy conditions.Getter, cannot patch", after.GetObjectKind())
	}

	// Store the diff from the before/after object, and return early if there are no changes.
	diff, err := conditions.NewPatch(
		before,
		after,
	)
	if err != nil {
		return errors.Wrapf(err, "object can not be patched")
	}
	if diff.IsZero() {
		return nil
	}

	// Make a copy of the object and store the key used if we have conflicts.
	key := client.ObjectKeyFromObject(after)

	// Define and start a backoff loop to handle conflicts
	// between controllers working on the same object.
	//
	// This has been copied from https://github.com/kubernetes/kubernetes/blob/release-1.16/pkg/controller/controller_utils.go#L86-L88.
	backoff := wait.Backoff{
		Steps:    5,
		Duration: 100 * time.Millisecond,
		Jitter:   1.0,
	}

	// Start the backoff loop and return errors if any.
	return wait.ExponentialBackoff(backoff, func() (bool, error) {
		latest, ok := before.DeepCopyObject().(conditions.Setter)
		if !ok {
			return false, errors.Errorf("object %s doesn't satisfy conditions.Setter, cannot patch", latest.GetObjectKind())
		}

		// Get a new copy of the object.
		if err := h.client.Get(ctx, key, latest); err != nil {
			return false, err
		}

		// Create the condition patch before merging conditions.
		conditionsPatch := client.MergeFromWithOptions(latest.DeepCopyObject().(conditions.Setter), client.MergeFromWithOptimisticLock{})

		// Set the condition patch previously created on the new object.
		if err := diff.Apply(latest, conditions.WithForceOverwrite(forceOverwrite), conditions.WithOwnedConditions(ownedConditions...)); err != nil {
			return false, err
		}

		// Issue the patch.
		err := h.client.Status().Patch(ctx, latest, conditionsPatch)
		switch {
		case apierrors.IsConflict(err):
			// Requeue.
			return false, nil
		case err != nil:
			return false, err
		default:
			return true, nil
		}
	})
}

// calculatePatch returns the before/after objects to be given in a controller-runtime patch, scoped down to the absolute necessary.
func (h *Helper) calculatePatch(afterObj client.Object, focus patchType) (client.Object, client.Object, error) {
	// Get a shallow unsafe copy of the before/after object in unstructured form.
	before := unsafeUnstructuredCopy(h.before, focus, h.isConditionsSetter)
	after := unsafeUnstructuredCopy(h.after, focus, h.isConditionsSetter)

	// We've now applied all modifications to local unstructured objects,
	// make copies of the original objects and convert them back.
	beforeObj := h.beforeObject.DeepCopyObject().(client.Object)
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(before.Object, beforeObj); err != nil {
		return nil, nil, err
	}
	afterObj = afterObj.DeepCopyObject().(client.Object)
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(after.Object, afterObj); err != nil {
		return nil, nil, err
	}
	return beforeObj, afterObj, nil
}

func (h *Helper) shouldPatch(in string) bool {
	return h.changes[in]
}

// calculate changes tries to build a patch from the before/after objects we have
// and store in a map which top-level fields (e.g. `metadata`, `spec`, `status`, etc.) have changed.
func (h *Helper) calculateChanges(after client.Object) (map[string]bool, error) {
	// Calculate patch data.
	patch := client.MergeFrom(h.beforeObject)
	diff, err := patch.Data(after)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to calculate patch data")
	}

	// Unmarshal patch data into a local map.
	patchDiff := map[string]any{}
	if err := json.Unmarshal(diff, &patchDiff); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal patch data into a map")
	}

	// Return the map.
	res := make(map[string]bool, len(patchDiff))
	for key := range patchDiff {
		res[key] = true
	}
	return res, nil
}
