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

// PatchedWritablePowerFeedRequestPhase * `single-phase` - Single phase * `three-phase` - Three-phase
type PatchedWritablePowerFeedRequestPhase string

// List of PatchedWritablePowerFeedRequest_phase
const (
	PATCHEDWRITABLEPOWERFEEDREQUESTPHASE_SINGLE_PHASE PatchedWritablePowerFeedRequestPhase = "single-phase"
	PATCHEDWRITABLEPOWERFEEDREQUESTPHASE_THREE_PHASE  PatchedWritablePowerFeedRequestPhase = "three-phase"
)

// All allowed values of PatchedWritablePowerFeedRequestPhase enum
var AllowedPatchedWritablePowerFeedRequestPhaseEnumValues = []PatchedWritablePowerFeedRequestPhase{
	"single-phase",
	"three-phase",
}

func (v *PatchedWritablePowerFeedRequestPhase) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritablePowerFeedRequestPhase(value)
	for _, existing := range AllowedPatchedWritablePowerFeedRequestPhaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritablePowerFeedRequestPhase", value)
}

// NewPatchedWritablePowerFeedRequestPhaseFromValue returns a pointer to a valid PatchedWritablePowerFeedRequestPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritablePowerFeedRequestPhaseFromValue(v string) (*PatchedWritablePowerFeedRequestPhase, error) {
	ev := PatchedWritablePowerFeedRequestPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritablePowerFeedRequestPhase: valid values are %v", v, AllowedPatchedWritablePowerFeedRequestPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritablePowerFeedRequestPhase) IsValid() bool {
	for _, existing := range AllowedPatchedWritablePowerFeedRequestPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritablePowerFeedRequest_phase value
func (v PatchedWritablePowerFeedRequestPhase) Ptr() *PatchedWritablePowerFeedRequestPhase {
	return &v
}

type NullablePatchedWritablePowerFeedRequestPhase struct {
	value *PatchedWritablePowerFeedRequestPhase
	isSet bool
}

func (v NullablePatchedWritablePowerFeedRequestPhase) Get() *PatchedWritablePowerFeedRequestPhase {
	return v.value
}

func (v *NullablePatchedWritablePowerFeedRequestPhase) Set(val *PatchedWritablePowerFeedRequestPhase) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerFeedRequestPhase) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerFeedRequestPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerFeedRequestPhase(val *PatchedWritablePowerFeedRequestPhase) *NullablePatchedWritablePowerFeedRequestPhase {
	return &NullablePatchedWritablePowerFeedRequestPhase{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerFeedRequestPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerFeedRequestPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
