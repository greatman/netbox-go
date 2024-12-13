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

// VLANQinqRoleLabel the model 'VLANQinqRoleLabel'
type VLANQinqRoleLabel string

// List of VLAN_qinq_role_label
const (
	VLANQINQROLELABEL_SERVICE  VLANQinqRoleLabel = "Service"
	VLANQINQROLELABEL_CUSTOMER VLANQinqRoleLabel = "Customer"
)

// All allowed values of VLANQinqRoleLabel enum
var AllowedVLANQinqRoleLabelEnumValues = []VLANQinqRoleLabel{
	"Service",
	"Customer",
}

func (v *VLANQinqRoleLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VLANQinqRoleLabel(value)
	for _, existing := range AllowedVLANQinqRoleLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VLANQinqRoleLabel", value)
}

// NewVLANQinqRoleLabelFromValue returns a pointer to a valid VLANQinqRoleLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVLANQinqRoleLabelFromValue(v string) (*VLANQinqRoleLabel, error) {
	ev := VLANQinqRoleLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VLANQinqRoleLabel: valid values are %v", v, AllowedVLANQinqRoleLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VLANQinqRoleLabel) IsValid() bool {
	for _, existing := range AllowedVLANQinqRoleLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VLAN_qinq_role_label value
func (v VLANQinqRoleLabel) Ptr() *VLANQinqRoleLabel {
	return &v
}

type NullableVLANQinqRoleLabel struct {
	value *VLANQinqRoleLabel
	isSet bool
}

func (v NullableVLANQinqRoleLabel) Get() *VLANQinqRoleLabel {
	return v.value
}

func (v *NullableVLANQinqRoleLabel) Set(val *VLANQinqRoleLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableVLANQinqRoleLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableVLANQinqRoleLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVLANQinqRoleLabel(val *VLANQinqRoleLabel) *NullableVLANQinqRoleLabel {
	return &NullableVLANQinqRoleLabel{value: val, isSet: true}
}

func (v NullableVLANQinqRoleLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVLANQinqRoleLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
