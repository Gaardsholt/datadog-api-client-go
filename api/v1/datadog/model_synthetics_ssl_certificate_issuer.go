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

// SyntheticsSslCertificateIssuer struct for SyntheticsSslCertificateIssuer
type SyntheticsSslCertificateIssuer struct {
	C  *string `json:"C,omitempty"`
	CN *string `json:"CN,omitempty"`
	L  *string `json:"L,omitempty"`
	O  *string `json:"O,omitempty"`
	OU *string `json:"OU,omitempty"`
	ST *string `json:"ST,omitempty"`
}

// NewSyntheticsSslCertificateIssuer instantiates a new SyntheticsSslCertificateIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsSslCertificateIssuer() *SyntheticsSslCertificateIssuer {
	this := SyntheticsSslCertificateIssuer{}
	return &this
}

// NewSyntheticsSslCertificateIssuerWithDefaults instantiates a new SyntheticsSslCertificateIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsSslCertificateIssuerWithDefaults() *SyntheticsSslCertificateIssuer {
	this := SyntheticsSslCertificateIssuer{}
	return &this
}

// GetC returns the C field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateIssuer) GetC() string {
	if o == nil || o.C == nil {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateIssuer) GetCOk() (string, bool) {
	if o == nil || o.C == nil {
		var ret string
		return ret, false
	}
	return *o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateIssuer) HasC() bool {
	if o != nil && o.C != nil {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *SyntheticsSslCertificateIssuer) SetC(v string) {
	o.C = &v
}

// GetCN returns the CN field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateIssuer) GetCN() string {
	if o == nil || o.CN == nil {
		var ret string
		return ret
	}
	return *o.CN
}

// GetCNOk returns a tuple with the CN field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateIssuer) GetCNOk() (string, bool) {
	if o == nil || o.CN == nil {
		var ret string
		return ret, false
	}
	return *o.CN, true
}

// HasCN returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateIssuer) HasCN() bool {
	if o != nil && o.CN != nil {
		return true
	}

	return false
}

// SetCN gets a reference to the given string and assigns it to the CN field.
func (o *SyntheticsSslCertificateIssuer) SetCN(v string) {
	o.CN = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateIssuer) GetL() string {
	if o == nil || o.L == nil {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateIssuer) GetLOk() (string, bool) {
	if o == nil || o.L == nil {
		var ret string
		return ret, false
	}
	return *o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateIssuer) HasL() bool {
	if o != nil && o.L != nil {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *SyntheticsSslCertificateIssuer) SetL(v string) {
	o.L = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateIssuer) GetO() string {
	if o == nil || o.O == nil {
		var ret string
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateIssuer) GetOOk() (string, bool) {
	if o == nil || o.O == nil {
		var ret string
		return ret, false
	}
	return *o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateIssuer) HasO() bool {
	if o != nil && o.O != nil {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *SyntheticsSslCertificateIssuer) SetO(v string) {
	o.O = &v
}

// GetOU returns the OU field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateIssuer) GetOU() string {
	if o == nil || o.OU == nil {
		var ret string
		return ret
	}
	return *o.OU
}

// GetOUOk returns a tuple with the OU field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateIssuer) GetOUOk() (string, bool) {
	if o == nil || o.OU == nil {
		var ret string
		return ret, false
	}
	return *o.OU, true
}

// HasOU returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateIssuer) HasOU() bool {
	if o != nil && o.OU != nil {
		return true
	}

	return false
}

// SetOU gets a reference to the given string and assigns it to the OU field.
func (o *SyntheticsSslCertificateIssuer) SetOU(v string) {
	o.OU = &v
}

// GetST returns the ST field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateIssuer) GetST() string {
	if o == nil || o.ST == nil {
		var ret string
		return ret
	}
	return *o.ST
}

// GetSTOk returns a tuple with the ST field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateIssuer) GetSTOk() (string, bool) {
	if o == nil || o.ST == nil {
		var ret string
		return ret, false
	}
	return *o.ST, true
}

// HasST returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateIssuer) HasST() bool {
	if o != nil && o.ST != nil {
		return true
	}

	return false
}

// SetST gets a reference to the given string and assigns it to the ST field.
func (o *SyntheticsSslCertificateIssuer) SetST(v string) {
	o.ST = &v
}

func (o SyntheticsSslCertificateIssuer) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if o.C != nil {
		toSerialize["C"] = o.C
	}
	if o.CN != nil {
		toSerialize["CN"] = o.CN
	}
	if o.L != nil {
		toSerialize["L"] = o.L
	}
	if o.O != nil {
		toSerialize["O"] = o.O
	}
	if o.OU != nil {
		toSerialize["OU"] = o.OU
	}
	if o.ST != nil {
		toSerialize["ST"] = o.ST
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsSslCertificateIssuer struct {
	value *SyntheticsSslCertificateIssuer
	isSet bool
}

func (v NullableSyntheticsSslCertificateIssuer) Get() *SyntheticsSslCertificateIssuer {
	return v.value
}

func (v NullableSyntheticsSslCertificateIssuer) Set(val *SyntheticsSslCertificateIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsSslCertificateIssuer) IsSet() bool {
	return v.isSet
}

func (v NullableSyntheticsSslCertificateIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsSslCertificateIssuer(val *SyntheticsSslCertificateIssuer) *NullableSyntheticsSslCertificateIssuer {
	return &NullableSyntheticsSslCertificateIssuer{value: val, isSet: true}
}

func (v NullableSyntheticsSslCertificateIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsSslCertificateIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}