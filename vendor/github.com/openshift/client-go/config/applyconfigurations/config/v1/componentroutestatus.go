// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ComponentRouteStatusApplyConfiguration represents a declarative configuration of the ComponentRouteStatus type for use
// with apply.
type ComponentRouteStatusApplyConfiguration struct {
	Namespace        *string                              `json:"namespace,omitempty"`
	Name             *string                              `json:"name,omitempty"`
	DefaultHostname  *configv1.Hostname                   `json:"defaultHostname,omitempty"`
	ConsumingUsers   []configv1.ConsumingUser             `json:"consumingUsers,omitempty"`
	CurrentHostnames []configv1.Hostname                  `json:"currentHostnames,omitempty"`
	Conditions       []metav1.ConditionApplyConfiguration `json:"conditions,omitempty"`
	RelatedObjects   []ObjectReferenceApplyConfiguration  `json:"relatedObjects,omitempty"`
}

// ComponentRouteStatusApplyConfiguration constructs a declarative configuration of the ComponentRouteStatus type for use with
// apply.
func ComponentRouteStatus() *ComponentRouteStatusApplyConfiguration {
	return &ComponentRouteStatusApplyConfiguration{}
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ComponentRouteStatusApplyConfiguration) WithNamespace(value string) *ComponentRouteStatusApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ComponentRouteStatusApplyConfiguration) WithName(value string) *ComponentRouteStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithDefaultHostname sets the DefaultHostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultHostname field is set to the value of the last call.
func (b *ComponentRouteStatusApplyConfiguration) WithDefaultHostname(value configv1.Hostname) *ComponentRouteStatusApplyConfiguration {
	b.DefaultHostname = &value
	return b
}

// WithConsumingUsers adds the given value to the ConsumingUsers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ConsumingUsers field.
func (b *ComponentRouteStatusApplyConfiguration) WithConsumingUsers(values ...configv1.ConsumingUser) *ComponentRouteStatusApplyConfiguration {
	for i := range values {
		b.ConsumingUsers = append(b.ConsumingUsers, values[i])
	}
	return b
}

// WithCurrentHostnames adds the given value to the CurrentHostnames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CurrentHostnames field.
func (b *ComponentRouteStatusApplyConfiguration) WithCurrentHostnames(values ...configv1.Hostname) *ComponentRouteStatusApplyConfiguration {
	for i := range values {
		b.CurrentHostnames = append(b.CurrentHostnames, values[i])
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ComponentRouteStatusApplyConfiguration) WithConditions(values ...*metav1.ConditionApplyConfiguration) *ComponentRouteStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithRelatedObjects adds the given value to the RelatedObjects field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RelatedObjects field.
func (b *ComponentRouteStatusApplyConfiguration) WithRelatedObjects(values ...*ObjectReferenceApplyConfiguration) *ComponentRouteStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRelatedObjects")
		}
		b.RelatedObjects = append(b.RelatedObjects, *values[i])
	}
	return b
}
