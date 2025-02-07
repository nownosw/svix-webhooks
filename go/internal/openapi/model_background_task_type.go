/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BackgroundTaskType the model 'BackgroundTaskType'
type BackgroundTaskType string

// List of BackgroundTaskType
const (
	BACKGROUNDTASKTYPE_ENDPOINT_REPLAY BackgroundTaskType = "endpoint.replay"
	BACKGROUNDTASKTYPE_ENDPOINT_RECOVER BackgroundTaskType = "endpoint.recover"
	BACKGROUNDTASKTYPE_APPLICATION_STATS BackgroundTaskType = "application.stats"
	BACKGROUNDTASKTYPE_MESSAGE_BROADCAST BackgroundTaskType = "message.broadcast"
	BACKGROUNDTASKTYPE_SDK_GENERATE BackgroundTaskType = "sdk.generate"
)

var allowedBackgroundTaskTypeEnumValues = []BackgroundTaskType{
	"endpoint.replay",
	"endpoint.recover",
	"application.stats",
	"message.broadcast",
	"sdk.generate",
}

func (v *BackgroundTaskType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackgroundTaskType(value)
	for _, existing := range allowedBackgroundTaskTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackgroundTaskType", value)
}

// NewBackgroundTaskTypeFromValue returns a pointer to a valid BackgroundTaskType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackgroundTaskTypeFromValue(v string) (*BackgroundTaskType, error) {
	ev := BackgroundTaskType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackgroundTaskType: valid values are %v", v, allowedBackgroundTaskTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackgroundTaskType) IsValid() bool {
	for _, existing := range allowedBackgroundTaskTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackgroundTaskType value
func (v BackgroundTaskType) Ptr() *BackgroundTaskType {
	return &v
}

type NullableBackgroundTaskType struct {
	value *BackgroundTaskType
	isSet bool
}

func (v NullableBackgroundTaskType) Get() *BackgroundTaskType {
	return v.value
}

func (v *NullableBackgroundTaskType) Set(val *BackgroundTaskType) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundTaskType) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundTaskType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundTaskType(val *BackgroundTaskType) *NullableBackgroundTaskType {
	return &NullableBackgroundTaskType{value: val, isSet: true}
}

func (v NullableBackgroundTaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundTaskType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

