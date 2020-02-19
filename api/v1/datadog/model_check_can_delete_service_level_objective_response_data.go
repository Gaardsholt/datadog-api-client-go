/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// CheckCanDeleteServiceLevelObjectiveResponseData An array of service level objective objects.
type CheckCanDeleteServiceLevelObjectiveResponseData struct {
	// An array of of SLO IDs that can be safely deleted.
	Ok *[]string `json:"ok,omitempty"`
}

// NewCheckCanDeleteServiceLevelObjectiveResponseData instantiates a new CheckCanDeleteServiceLevelObjectiveResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckCanDeleteServiceLevelObjectiveResponseData() *CheckCanDeleteServiceLevelObjectiveResponseData {
	this := CheckCanDeleteServiceLevelObjectiveResponseData{}
	return &this
}

// NewCheckCanDeleteServiceLevelObjectiveResponseDataWithDefaults instantiates a new CheckCanDeleteServiceLevelObjectiveResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckCanDeleteServiceLevelObjectiveResponseDataWithDefaults() *CheckCanDeleteServiceLevelObjectiveResponseData {
	this := CheckCanDeleteServiceLevelObjectiveResponseData{}
	return &this
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *CheckCanDeleteServiceLevelObjectiveResponseData) GetOk() []string {
	if o == nil || o.Ok == nil {
		var ret []string
		return ret
	}
	return *o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CheckCanDeleteServiceLevelObjectiveResponseData) GetOkOk() ([]string, bool) {
	if o == nil || o.Ok == nil {
		var ret []string
		return ret, false
	}
	return *o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *CheckCanDeleteServiceLevelObjectiveResponseData) HasOk() bool {
	if o != nil && o.Ok != nil {
		return true
	}

	return false
}

// SetOk gets a reference to the given []string and assigns it to the Ok field.
func (o *CheckCanDeleteServiceLevelObjectiveResponseData) SetOk(v []string) {
	o.Ok = &v
}

func (o CheckCanDeleteServiceLevelObjectiveResponseData) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if o.Ok != nil {
		toSerialize["ok"] = o.Ok
	}
	return json.Marshal(toSerialize)
}

type NullableCheckCanDeleteServiceLevelObjectiveResponseData struct {
	value *CheckCanDeleteServiceLevelObjectiveResponseData
	isSet bool
}

func (v NullableCheckCanDeleteServiceLevelObjectiveResponseData) Get() *CheckCanDeleteServiceLevelObjectiveResponseData {
	return v.value
}

func (v NullableCheckCanDeleteServiceLevelObjectiveResponseData) Set(val *CheckCanDeleteServiceLevelObjectiveResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckCanDeleteServiceLevelObjectiveResponseData) IsSet() bool {
	return v.isSet
}

func (v NullableCheckCanDeleteServiceLevelObjectiveResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckCanDeleteServiceLevelObjectiveResponseData(val *CheckCanDeleteServiceLevelObjectiveResponseData) *NullableCheckCanDeleteServiceLevelObjectiveResponseData {
	return &NullableCheckCanDeleteServiceLevelObjectiveResponseData{value: val, isSet: true}
}

func (v NullableCheckCanDeleteServiceLevelObjectiveResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckCanDeleteServiceLevelObjectiveResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
