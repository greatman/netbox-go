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

// checks if the InterfaceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceType{}

// InterfaceType struct for InterfaceType
type InterfaceType struct {
	Value                *InterfaceTypeValue `json:"value,omitempty"`
	Label                *InterfaceTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceType InterfaceType

// NewInterfaceType instantiates a new InterfaceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceType() *InterfaceType {
	this := InterfaceType{}
	return &this
}

// NewInterfaceTypeWithDefaults instantiates a new InterfaceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTypeWithDefaults() *InterfaceType {
	this := InterfaceType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfaceType) GetValue() InterfaceTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret InterfaceTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceType) GetValueOk() (*InterfaceTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given InterfaceTypeValue and assigns it to the Value field.
func (o *InterfaceType) SetValue(v InterfaceTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceType) GetLabel() InterfaceTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret InterfaceTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceType) GetLabelOk() (*InterfaceTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given InterfaceTypeLabel and assigns it to the Label field.
func (o *InterfaceType) SetLabel(v InterfaceTypeLabel) {
	o.Label = &v
}

func (o InterfaceType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceType) ToMap() (map[string]interface{}, error) {
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

func (o *InterfaceType) UnmarshalJSON(data []byte) (err error) {
	varInterfaceType := _InterfaceType{}

	err = json.Unmarshal(data, &varInterfaceType)

	if err != nil {
		return err
	}

	*o = InterfaceType(varInterfaceType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceType struct {
	value *InterfaceType
	isSet bool
}

func (v NullableInterfaceType) Get() *InterfaceType {
	return v.value
}

func (v *NullableInterfaceType) Set(val *InterfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceType(val *InterfaceType) *NullableInterfaceType {
	return &NullableInterfaceType{value: val, isSet: true}
}

func (v NullableInterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}