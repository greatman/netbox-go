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

// checks if the ObjectChangeAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectChangeAction{}

// ObjectChangeAction struct for ObjectChangeAction
type ObjectChangeAction struct {
	Value                *ObjectChangeActionValue `json:"value,omitempty"`
	Label                *ObjectChangeActionLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ObjectChangeAction ObjectChangeAction

// NewObjectChangeAction instantiates a new ObjectChangeAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectChangeAction() *ObjectChangeAction {
	this := ObjectChangeAction{}
	return &this
}

// NewObjectChangeActionWithDefaults instantiates a new ObjectChangeAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectChangeActionWithDefaults() *ObjectChangeAction {
	this := ObjectChangeAction{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ObjectChangeAction) GetValue() ObjectChangeActionValue {
	if o == nil || IsNil(o.Value) {
		var ret ObjectChangeActionValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectChangeAction) GetValueOk() (*ObjectChangeActionValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ObjectChangeAction) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ObjectChangeActionValue and assigns it to the Value field.
func (o *ObjectChangeAction) SetValue(v ObjectChangeActionValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ObjectChangeAction) GetLabel() ObjectChangeActionLabel {
	if o == nil || IsNil(o.Label) {
		var ret ObjectChangeActionLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectChangeAction) GetLabelOk() (*ObjectChangeActionLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ObjectChangeAction) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given ObjectChangeActionLabel and assigns it to the Label field.
func (o *ObjectChangeAction) SetLabel(v ObjectChangeActionLabel) {
	o.Label = &v
}

func (o ObjectChangeAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectChangeAction) ToMap() (map[string]interface{}, error) {
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

func (o *ObjectChangeAction) UnmarshalJSON(data []byte) (err error) {
	varObjectChangeAction := _ObjectChangeAction{}

	err = json.Unmarshal(data, &varObjectChangeAction)

	if err != nil {
		return err
	}

	*o = ObjectChangeAction(varObjectChangeAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjectChangeAction struct {
	value *ObjectChangeAction
	isSet bool
}

func (v NullableObjectChangeAction) Get() *ObjectChangeAction {
	return v.value
}

func (v *NullableObjectChangeAction) Set(val *ObjectChangeAction) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectChangeAction) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectChangeAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectChangeAction(val *ObjectChangeAction) *NullableObjectChangeAction {
	return &NullableObjectChangeAction{value: val, isSet: true}
}

func (v NullableObjectChangeAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectChangeAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
