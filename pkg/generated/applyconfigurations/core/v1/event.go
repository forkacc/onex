// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	internal "github.com/onexstack/onex/pkg/generated/applyconfigurations/internal"
	metav1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/meta/v1"
	corev1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// EventApplyConfiguration represents a declarative configuration of the Event type for use
// with apply.
type EventApplyConfiguration struct {
	metav1.TypeMetaApplyConfiguration    `json:",inline"`
	*metav1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	InvolvedObject                       *ObjectReferenceApplyConfiguration `json:"involvedObject,omitempty"`
	Reason                               *string                            `json:"reason,omitempty"`
	Message                              *string                            `json:"message,omitempty"`
	Source                               *EventSourceApplyConfiguration     `json:"source,omitempty"`
	FirstTimestamp                       *apismetav1.Time                   `json:"firstTimestamp,omitempty"`
	LastTimestamp                        *apismetav1.Time                   `json:"lastTimestamp,omitempty"`
	Count                                *int32                             `json:"count,omitempty"`
	Type                                 *string                            `json:"type,omitempty"`
	EventTime                            *apismetav1.MicroTime              `json:"eventTime,omitempty"`
	Series                               *EventSeriesApplyConfiguration     `json:"series,omitempty"`
	Action                               *string                            `json:"action,omitempty"`
	Related                              *ObjectReferenceApplyConfiguration `json:"related,omitempty"`
	ReportingController                  *string                            `json:"reportingComponent,omitempty"`
	ReportingInstance                    *string                            `json:"reportingInstance,omitempty"`
}

// Event constructs a declarative configuration of the Event type for use with
// apply.
func Event(name, namespace string) *EventApplyConfiguration {
	b := &EventApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("Event")
	b.WithAPIVersion("v1")
	return b
}

// ExtractEvent extracts the applied configuration owned by fieldManager from
// event. If no managedFields are found in event for fieldManager, a
// EventApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// event must be a unmodified Event API object that was retrieved from the Kubernetes API.
// ExtractEvent provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractEvent(event *corev1.Event, fieldManager string) (*EventApplyConfiguration, error) {
	return extractEvent(event, fieldManager, "")
}

// ExtractEventStatus is the same as ExtractEvent except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractEventStatus(event *corev1.Event, fieldManager string) (*EventApplyConfiguration, error) {
	return extractEvent(event, fieldManager, "status")
}

func extractEvent(event *corev1.Event, fieldManager string, subresource string) (*EventApplyConfiguration, error) {
	b := &EventApplyConfiguration{}
	err := managedfields.ExtractInto(event, internal.Parser().Type("io.k8s.api.core.v1.Event"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(event.Name)
	b.WithNamespace(event.Namespace)

	b.WithKind("Event")
	b.WithAPIVersion("v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *EventApplyConfiguration) WithKind(value string) *EventApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *EventApplyConfiguration) WithAPIVersion(value string) *EventApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *EventApplyConfiguration) WithName(value string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *EventApplyConfiguration) WithGenerateName(value string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *EventApplyConfiguration) WithNamespace(value string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *EventApplyConfiguration) WithUID(value types.UID) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *EventApplyConfiguration) WithResourceVersion(value string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *EventApplyConfiguration) WithGeneration(value int64) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *EventApplyConfiguration) WithCreationTimestamp(value apismetav1.Time) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *EventApplyConfiguration) WithDeletionTimestamp(value apismetav1.Time) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *EventApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *EventApplyConfiguration) WithLabels(entries map[string]string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *EventApplyConfiguration) WithAnnotations(entries map[string]string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *EventApplyConfiguration) WithOwnerReferences(values ...*metav1.OwnerReferenceApplyConfiguration) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *EventApplyConfiguration) WithFinalizers(values ...string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *EventApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &metav1.ObjectMetaApplyConfiguration{}
	}
}

// WithInvolvedObject sets the InvolvedObject field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InvolvedObject field is set to the value of the last call.
func (b *EventApplyConfiguration) WithInvolvedObject(value *ObjectReferenceApplyConfiguration) *EventApplyConfiguration {
	b.InvolvedObject = value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *EventApplyConfiguration) WithReason(value string) *EventApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *EventApplyConfiguration) WithMessage(value string) *EventApplyConfiguration {
	b.Message = &value
	return b
}

// WithSource sets the Source field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Source field is set to the value of the last call.
func (b *EventApplyConfiguration) WithSource(value *EventSourceApplyConfiguration) *EventApplyConfiguration {
	b.Source = value
	return b
}

// WithFirstTimestamp sets the FirstTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FirstTimestamp field is set to the value of the last call.
func (b *EventApplyConfiguration) WithFirstTimestamp(value apismetav1.Time) *EventApplyConfiguration {
	b.FirstTimestamp = &value
	return b
}

// WithLastTimestamp sets the LastTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTimestamp field is set to the value of the last call.
func (b *EventApplyConfiguration) WithLastTimestamp(value apismetav1.Time) *EventApplyConfiguration {
	b.LastTimestamp = &value
	return b
}

// WithCount sets the Count field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Count field is set to the value of the last call.
func (b *EventApplyConfiguration) WithCount(value int32) *EventApplyConfiguration {
	b.Count = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *EventApplyConfiguration) WithType(value string) *EventApplyConfiguration {
	b.Type = &value
	return b
}

// WithEventTime sets the EventTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EventTime field is set to the value of the last call.
func (b *EventApplyConfiguration) WithEventTime(value apismetav1.MicroTime) *EventApplyConfiguration {
	b.EventTime = &value
	return b
}

// WithSeries sets the Series field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Series field is set to the value of the last call.
func (b *EventApplyConfiguration) WithSeries(value *EventSeriesApplyConfiguration) *EventApplyConfiguration {
	b.Series = value
	return b
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *EventApplyConfiguration) WithAction(value string) *EventApplyConfiguration {
	b.Action = &value
	return b
}

// WithRelated sets the Related field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Related field is set to the value of the last call.
func (b *EventApplyConfiguration) WithRelated(value *ObjectReferenceApplyConfiguration) *EventApplyConfiguration {
	b.Related = value
	return b
}

// WithReportingController sets the ReportingController field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReportingController field is set to the value of the last call.
func (b *EventApplyConfiguration) WithReportingController(value string) *EventApplyConfiguration {
	b.ReportingController = &value
	return b
}

// WithReportingInstance sets the ReportingInstance field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReportingInstance field is set to the value of the last call.
func (b *EventApplyConfiguration) WithReportingInstance(value string) *EventApplyConfiguration {
	b.ReportingInstance = &value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *EventApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}
