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

// PatchedWritablePowerOutletRequestFeedLeg Phase (for three-phase feeds)  * `A` - A * `B` - B * `C` - C
type PatchedWritablePowerOutletRequestFeedLeg string

// List of PatchedWritablePowerOutletRequest_feed_leg
const (
	PATCHEDWRITABLEPOWEROUTLETREQUESTFEEDLEG_A     PatchedWritablePowerOutletRequestFeedLeg = "A"
	PATCHEDWRITABLEPOWEROUTLETREQUESTFEEDLEG_B     PatchedWritablePowerOutletRequestFeedLeg = "B"
	PATCHEDWRITABLEPOWEROUTLETREQUESTFEEDLEG_C     PatchedWritablePowerOutletRequestFeedLeg = "C"
	PATCHEDWRITABLEPOWEROUTLETREQUESTFEEDLEG_EMPTY PatchedWritablePowerOutletRequestFeedLeg = ""
)

// All allowed values of PatchedWritablePowerOutletRequestFeedLeg enum
var AllowedPatchedWritablePowerOutletRequestFeedLegEnumValues = []PatchedWritablePowerOutletRequestFeedLeg{
	"A",
	"B",
	"C",
	"",
}

func (v *PatchedWritablePowerOutletRequestFeedLeg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritablePowerOutletRequestFeedLeg(value)
	for _, existing := range AllowedPatchedWritablePowerOutletRequestFeedLegEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritablePowerOutletRequestFeedLeg", value)
}

// NewPatchedWritablePowerOutletRequestFeedLegFromValue returns a pointer to a valid PatchedWritablePowerOutletRequestFeedLeg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritablePowerOutletRequestFeedLegFromValue(v string) (*PatchedWritablePowerOutletRequestFeedLeg, error) {
	ev := PatchedWritablePowerOutletRequestFeedLeg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritablePowerOutletRequestFeedLeg: valid values are %v", v, AllowedPatchedWritablePowerOutletRequestFeedLegEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritablePowerOutletRequestFeedLeg) IsValid() bool {
	for _, existing := range AllowedPatchedWritablePowerOutletRequestFeedLegEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritablePowerOutletRequest_feed_leg value
func (v PatchedWritablePowerOutletRequestFeedLeg) Ptr() *PatchedWritablePowerOutletRequestFeedLeg {
	return &v
}

type NullablePatchedWritablePowerOutletRequestFeedLeg struct {
	value *PatchedWritablePowerOutletRequestFeedLeg
	isSet bool
}

func (v NullablePatchedWritablePowerOutletRequestFeedLeg) Get() *PatchedWritablePowerOutletRequestFeedLeg {
	return v.value
}

func (v *NullablePatchedWritablePowerOutletRequestFeedLeg) Set(val *PatchedWritablePowerOutletRequestFeedLeg) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerOutletRequestFeedLeg) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerOutletRequestFeedLeg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerOutletRequestFeedLeg(val *PatchedWritablePowerOutletRequestFeedLeg) *NullablePatchedWritablePowerOutletRequestFeedLeg {
	return &NullablePatchedWritablePowerOutletRequestFeedLeg{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerOutletRequestFeedLeg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerOutletRequestFeedLeg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
