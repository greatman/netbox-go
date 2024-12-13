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

// CustomFieldChoiceSetBaseChoicesValue * `IATA` - IATA (Airport codes) * `ISO_3166` - ISO 3166 (Country codes) * `UN_LOCODE` - UN/LOCODE (Location codes)
type CustomFieldChoiceSetBaseChoicesValue string

// List of CustomFieldChoiceSet_base_choices_value
const (
	CUSTOMFIELDCHOICESETBASECHOICESVALUE_IATA      CustomFieldChoiceSetBaseChoicesValue = "IATA"
	CUSTOMFIELDCHOICESETBASECHOICESVALUE_ISO_3166  CustomFieldChoiceSetBaseChoicesValue = "ISO_3166"
	CUSTOMFIELDCHOICESETBASECHOICESVALUE_UN_LOCODE CustomFieldChoiceSetBaseChoicesValue = "UN_LOCODE"
)

// All allowed values of CustomFieldChoiceSetBaseChoicesValue enum
var AllowedCustomFieldChoiceSetBaseChoicesValueEnumValues = []CustomFieldChoiceSetBaseChoicesValue{
	"IATA",
	"ISO_3166",
	"UN_LOCODE",
}

func (v *CustomFieldChoiceSetBaseChoicesValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldChoiceSetBaseChoicesValue(value)
	for _, existing := range AllowedCustomFieldChoiceSetBaseChoicesValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldChoiceSetBaseChoicesValue", value)
}

// NewCustomFieldChoiceSetBaseChoicesValueFromValue returns a pointer to a valid CustomFieldChoiceSetBaseChoicesValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldChoiceSetBaseChoicesValueFromValue(v string) (*CustomFieldChoiceSetBaseChoicesValue, error) {
	ev := CustomFieldChoiceSetBaseChoicesValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldChoiceSetBaseChoicesValue: valid values are %v", v, AllowedCustomFieldChoiceSetBaseChoicesValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldChoiceSetBaseChoicesValue) IsValid() bool {
	for _, existing := range AllowedCustomFieldChoiceSetBaseChoicesValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomFieldChoiceSet_base_choices_value value
func (v CustomFieldChoiceSetBaseChoicesValue) Ptr() *CustomFieldChoiceSetBaseChoicesValue {
	return &v
}

type NullableCustomFieldChoiceSetBaseChoicesValue struct {
	value *CustomFieldChoiceSetBaseChoicesValue
	isSet bool
}

func (v NullableCustomFieldChoiceSetBaseChoicesValue) Get() *CustomFieldChoiceSetBaseChoicesValue {
	return v.value
}

func (v *NullableCustomFieldChoiceSetBaseChoicesValue) Set(val *CustomFieldChoiceSetBaseChoicesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldChoiceSetBaseChoicesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldChoiceSetBaseChoicesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldChoiceSetBaseChoicesValue(val *CustomFieldChoiceSetBaseChoicesValue) *NullableCustomFieldChoiceSetBaseChoicesValue {
	return &NullableCustomFieldChoiceSetBaseChoicesValue{value: val, isSet: true}
}

func (v NullableCustomFieldChoiceSetBaseChoicesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldChoiceSetBaseChoicesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}