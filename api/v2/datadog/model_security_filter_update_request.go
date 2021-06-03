/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SecurityFilterUpdateRequest The new security filter body.
type SecurityFilterUpdateRequest struct {
	Data SecurityFilterUpdateData `json:"data"`
}

// NewSecurityFilterUpdateRequest instantiates a new SecurityFilterUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityFilterUpdateRequest(data SecurityFilterUpdateData) *SecurityFilterUpdateRequest {
	this := SecurityFilterUpdateRequest{}
	this.Data = data
	return &this
}

// NewSecurityFilterUpdateRequestWithDefaults instantiates a new SecurityFilterUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityFilterUpdateRequestWithDefaults() *SecurityFilterUpdateRequest {
	this := SecurityFilterUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SecurityFilterUpdateRequest) GetData() SecurityFilterUpdateData {
	if o == nil {
		var ret SecurityFilterUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterUpdateRequest) GetDataOk() (*SecurityFilterUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SecurityFilterUpdateRequest) SetData(v SecurityFilterUpdateData) {
	o.Data = v
}

func (o SecurityFilterUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityFilterUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Data *SecurityFilterUpdateData `json:"data"`
	}{}
	all := struct {
		Data SecurityFilterUpdateData `json:"data"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("Required field data missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Data = all.Data
	return nil
}

type NullableSecurityFilterUpdateRequest struct {
	value *SecurityFilterUpdateRequest
	isSet bool
}

func (v NullableSecurityFilterUpdateRequest) Get() *SecurityFilterUpdateRequest {
	return v.value
}

func (v *NullableSecurityFilterUpdateRequest) Set(val *SecurityFilterUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityFilterUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityFilterUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityFilterUpdateRequest(val *SecurityFilterUpdateRequest) *NullableSecurityFilterUpdateRequest {
	return &NullableSecurityFilterUpdateRequest{value: val, isSet: true}
}

func (v NullableSecurityFilterUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityFilterUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}