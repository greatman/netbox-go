/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.2-beta1 (4.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// ParentChildStatus1 Parent devices house child devices in device bays. Leave blank if this device type is neither a parent nor a child.  * `parent` - Parent * `child` - Child
type ParentChildStatus1 string

// List of Parent_child_status_1
const (
	PARENTCHILDSTATUS1_PARENT ParentChildStatus1 = "parent"
	PARENTCHILDSTATUS1_CHILD  ParentChildStatus1 = "child"
	PARENTCHILDSTATUS1_EMPTY  ParentChildStatus1 = ""
)

// All allowed values of ParentChildStatus1 enum
var AllowedParentChildStatus1EnumValues = []ParentChildStatus1{
	"parent",
	"child",
	"",
}

func (v *ParentChildStatus1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ParentChildStatus1(value)
	for _, existing := range AllowedParentChildStatus1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ParentChildStatus1", value)
}

// NewParentChildStatus1FromValue returns a pointer to a valid ParentChildStatus1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewParentChildStatus1FromValue(v string) (*ParentChildStatus1, error) {
	ev := ParentChildStatus1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ParentChildStatus1: valid values are %v", v, AllowedParentChildStatus1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ParentChildStatus1) IsValid() bool {
	for _, existing := range AllowedParentChildStatus1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Parent_child_status_1 value
func (v ParentChildStatus1) Ptr() *ParentChildStatus1 {
	return &v
}

type NullableParentChildStatus1 struct {
	value *ParentChildStatus1
	isSet bool
}

func (v NullableParentChildStatus1) Get() *ParentChildStatus1 {
	return v.value
}

func (v *NullableParentChildStatus1) Set(val *ParentChildStatus1) {
	v.value = val
	v.isSet = true
}

func (v NullableParentChildStatus1) IsSet() bool {
	return v.isSet
}

func (v *NullableParentChildStatus1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentChildStatus1(val *ParentChildStatus1) *NullableParentChildStatus1 {
	return &NullableParentChildStatus1{value: val, isSet: true}
}

func (v NullableParentChildStatus1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentChildStatus1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}