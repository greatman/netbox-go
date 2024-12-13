/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.2-beta1 (4.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the FrontPortType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrontPortType{}

// FrontPortType struct for FrontPortType
type FrontPortType struct {
	Value                *FrontPortTypeValue `json:"value,omitempty"`
	Label                *FrontPortTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FrontPortType FrontPortType

// NewFrontPortType instantiates a new FrontPortType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrontPortType() *FrontPortType {
	this := FrontPortType{}
	return &this
}

// NewFrontPortTypeWithDefaults instantiates a new FrontPortType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrontPortTypeWithDefaults() *FrontPortType {
	this := FrontPortType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FrontPortType) GetValue() FrontPortTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret FrontPortTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortType) GetValueOk() (*FrontPortTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FrontPortType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given FrontPortTypeValue and assigns it to the Value field.
func (o *FrontPortType) SetValue(v FrontPortTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FrontPortType) GetLabel() FrontPortTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret FrontPortTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortType) GetLabelOk() (*FrontPortTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FrontPortType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given FrontPortTypeLabel and assigns it to the Label field.
func (o *FrontPortType) SetLabel(v FrontPortTypeLabel) {
	o.Label = &v
}

func (o FrontPortType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrontPortType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrontPortType) UnmarshalJSON(data []byte) (err error) {
	varFrontPortType := _FrontPortType{}

	err = json.Unmarshal(data, &varFrontPortType)

	if err != nil {
		return err
	}

	*o = FrontPortType(varFrontPortType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFrontPortType struct {
	value *FrontPortType
	isSet bool
}

func (v NullableFrontPortType) Get() *FrontPortType {
	return v.value
}

func (v *NullableFrontPortType) Set(val *FrontPortType) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPortType) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPortType(val *FrontPortType) *NullableFrontPortType {
	return &NullableFrontPortType{value: val, isSet: true}
}

func (v NullableFrontPortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
