/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// DeletedMonitor Response from the DeleteMonitor call
type DeletedMonitor struct {
	// ID of the deleted monitor
	DeletedMonitorId *int64 `json:"deleted_monitor_id,omitempty"`
}

// GetDeletedMonitorId returns the DeletedMonitorId field value if set, zero value otherwise.
func (o *DeletedMonitor) GetDeletedMonitorId() int64 {
	if o == nil || o.DeletedMonitorId == nil {
		var ret int64
		return ret
	}
	return *o.DeletedMonitorId
}

// GetDeletedMonitorIdOk returns a tuple with the DeletedMonitorId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeletedMonitor) GetDeletedMonitorIdOk() (int64, bool) {
	if o == nil || o.DeletedMonitorId == nil {
		var ret int64
		return ret, false
	}
	return *o.DeletedMonitorId, true
}

// HasDeletedMonitorId returns a boolean if a field has been set.
func (o *DeletedMonitor) HasDeletedMonitorId() bool {
	if o != nil && o.DeletedMonitorId != nil {
		return true
	}

	return false
}

// SetDeletedMonitorId gets a reference to the given int64 and assigns it to the DeletedMonitorId field.
func (o *DeletedMonitor) SetDeletedMonitorId(v int64) {
	o.DeletedMonitorId = &v
}

type NullableDeletedMonitor struct {
	Value        DeletedMonitor
	ExplicitNull bool
}

func (v NullableDeletedMonitor) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableDeletedMonitor) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}