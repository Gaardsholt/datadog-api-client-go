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

// UserDisableResponse struct for UserDisableResponse
type UserDisableResponse struct {
	Message *string `json:"message,omitempty"`
}

// NewUserDisableResponse instantiates a new UserDisableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDisableResponse() *UserDisableResponse {
	this := UserDisableResponse{}
	return &this
}

// NewUserDisableResponseWithDefaults instantiates a new UserDisableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDisableResponseWithDefaults() *UserDisableResponse {
	this := UserDisableResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UserDisableResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UserDisableResponse) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UserDisableResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UserDisableResponse) SetMessage(v string) {
	o.Message = &v
}

func (o UserDisableResponse) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableUserDisableResponse struct {
	value *UserDisableResponse
	isSet bool
}

func (v NullableUserDisableResponse) Get() *UserDisableResponse {
	return v.value
}

func (v NullableUserDisableResponse) Set(val *UserDisableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDisableResponse) IsSet() bool {
	return v.isSet
}

func (v NullableUserDisableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDisableResponse(val *UserDisableResponse) *NullableUserDisableResponse {
	return &NullableUserDisableResponse{value: val, isSet: true}
}

func (v NullableUserDisableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDisableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
