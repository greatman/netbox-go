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

// WirelessLANAuthTypeValue * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise
type WirelessLANAuthTypeValue string

// List of WirelessLAN_auth_type_value
const (
	WIRELESSLANAUTHTYPEVALUE_OPEN           WirelessLANAuthTypeValue = "open"
	WIRELESSLANAUTHTYPEVALUE_WEP            WirelessLANAuthTypeValue = "wep"
	WIRELESSLANAUTHTYPEVALUE_WPA_PERSONAL   WirelessLANAuthTypeValue = "wpa-personal"
	WIRELESSLANAUTHTYPEVALUE_WPA_ENTERPRISE WirelessLANAuthTypeValue = "wpa-enterprise"
	WIRELESSLANAUTHTYPEVALUE_EMPTY          WirelessLANAuthTypeValue = ""
)

// All allowed values of WirelessLANAuthTypeValue enum
var AllowedWirelessLANAuthTypeValueEnumValues = []WirelessLANAuthTypeValue{
	"open",
	"wep",
	"wpa-personal",
	"wpa-enterprise",
	"",
}

func (v *WirelessLANAuthTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessLANAuthTypeValue(value)
	for _, existing := range AllowedWirelessLANAuthTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessLANAuthTypeValue", value)
}

// NewWirelessLANAuthTypeValueFromValue returns a pointer to a valid WirelessLANAuthTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessLANAuthTypeValueFromValue(v string) (*WirelessLANAuthTypeValue, error) {
	ev := WirelessLANAuthTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessLANAuthTypeValue: valid values are %v", v, AllowedWirelessLANAuthTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessLANAuthTypeValue) IsValid() bool {
	for _, existing := range AllowedWirelessLANAuthTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WirelessLAN_auth_type_value value
func (v WirelessLANAuthTypeValue) Ptr() *WirelessLANAuthTypeValue {
	return &v
}

type NullableWirelessLANAuthTypeValue struct {
	value *WirelessLANAuthTypeValue
	isSet bool
}

func (v NullableWirelessLANAuthTypeValue) Get() *WirelessLANAuthTypeValue {
	return v.value
}

func (v *NullableWirelessLANAuthTypeValue) Set(val *WirelessLANAuthTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLANAuthTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLANAuthTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLANAuthTypeValue(val *WirelessLANAuthTypeValue) *NullableWirelessLANAuthTypeValue {
	return &NullableWirelessLANAuthTypeValue{value: val, isSet: true}
}

func (v NullableWirelessLANAuthTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLANAuthTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
