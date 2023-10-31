/*
SimKube API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// checks if the ExportFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportFilters{}

//+kubebuilder:object:generate=false

// ExportFilters struct for ExportFilters
type ExportFilters struct {
	ExcludedNamespaces []string               `json:"excluded_namespaces"`
	ExcludedLabels     []metav1.LabelSelector `json:"excluded_labels"`
	ExcludeDaemonsets  bool                   `json:"exclude_daemonsets"`
}

// NewExportFilters instantiates a new ExportFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportFilters(excludedNamespaces []string, excludedLabels []metav1.LabelSelector, excludeDaemonsets bool) *ExportFilters {
	this := ExportFilters{}
	this.ExcludedNamespaces = excludedNamespaces
	this.ExcludedLabels = excludedLabels
	this.ExcludeDaemonsets = excludeDaemonsets
	return &this
}

// NewExportFiltersWithDefaults instantiates a new ExportFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportFiltersWithDefaults() *ExportFilters {
	this := ExportFilters{}
	return &this
}

// GetExcludedNamespaces returns the ExcludedNamespaces field value
func (o *ExportFilters) GetExcludedNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExcludedNamespaces
}

// GetExcludedNamespacesOk returns a tuple with the ExcludedNamespaces field value
// and a boolean to check if the value has been set.
func (o *ExportFilters) GetExcludedNamespacesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExcludedNamespaces, true
}

// SetExcludedNamespaces sets field value
func (o *ExportFilters) SetExcludedNamespaces(v []string) {
	o.ExcludedNamespaces = v
}

// GetExcludedLabels returns the ExcludedLabels field value
func (o *ExportFilters) GetExcludedLabels() []metav1.LabelSelector {
	if o == nil {
		var ret []metav1.LabelSelector
		return ret
	}

	return o.ExcludedLabels
}

// GetExcludedLabelsOk returns a tuple with the ExcludedLabels field value
// and a boolean to check if the value has been set.
func (o *ExportFilters) GetExcludedLabelsOk() ([]metav1.LabelSelector, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExcludedLabels, true
}

// SetExcludedLabels sets field value
func (o *ExportFilters) SetExcludedLabels(v []metav1.LabelSelector) {
	o.ExcludedLabels = v
}

// GetExcludeDaemonsets returns the ExcludeDaemonsets field value
func (o *ExportFilters) GetExcludeDaemonsets() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExcludeDaemonsets
}

// GetExcludeDaemonsetsOk returns a tuple with the ExcludeDaemonsets field value
// and a boolean to check if the value has been set.
func (o *ExportFilters) GetExcludeDaemonsetsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExcludeDaemonsets, true
}

// SetExcludeDaemonsets sets field value
func (o *ExportFilters) SetExcludeDaemonsets(v bool) {
	o.ExcludeDaemonsets = v
}

func (o ExportFilters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["excluded_namespaces"] = o.ExcludedNamespaces
	toSerialize["excluded_labels"] = o.ExcludedLabels
	toSerialize["exclude_daemonsets"] = o.ExcludeDaemonsets
	return toSerialize, nil
}

//+kubebuilder:object:generate=false

type NullableExportFilters struct {
	value *ExportFilters
	isSet bool
}

func (v NullableExportFilters) Get() *ExportFilters {
	return v.value
}

func (v *NullableExportFilters) Set(val *ExportFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableExportFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableExportFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportFilters(val *ExportFilters) *NullableExportFilters {
	return &NullableExportFilters{value: val, isSet: true}
}

func (v NullableExportFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
