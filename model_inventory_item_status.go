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

// checks if the InventoryItemStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryItemStatus{}

// InventoryItemStatus struct for InventoryItemStatus
type InventoryItemStatus struct {
	Value                *InventoryItemStatusValue `json:"value,omitempty"`
	Label                *InventoryItemStatusLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryItemStatus InventoryItemStatus

// NewInventoryItemStatus instantiates a new InventoryItemStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryItemStatus() *InventoryItemStatus {
	this := InventoryItemStatus{}
	return &this
}

// NewInventoryItemStatusWithDefaults instantiates a new InventoryItemStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryItemStatusWithDefaults() *InventoryItemStatus {
	this := InventoryItemStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InventoryItemStatus) GetValue() InventoryItemStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret InventoryItemStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemStatus) GetValueOk() (*InventoryItemStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InventoryItemStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given InventoryItemStatusValue and assigns it to the Value field.
func (o *InventoryItemStatus) SetValue(v InventoryItemStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InventoryItemStatus) GetLabel() InventoryItemStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret InventoryItemStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemStatus) GetLabelOk() (*InventoryItemStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InventoryItemStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given InventoryItemStatusLabel and assigns it to the Label field.
func (o *InventoryItemStatus) SetLabel(v InventoryItemStatusLabel) {
	o.Label = &v
}

func (o InventoryItemStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryItemStatus) ToMap() (map[string]interface{}, error) {
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

func (o *InventoryItemStatus) UnmarshalJSON(data []byte) (err error) {
	varInventoryItemStatus := _InventoryItemStatus{}

	err = json.Unmarshal(data, &varInventoryItemStatus)

	if err != nil {
		return err
	}

	*o = InventoryItemStatus(varInventoryItemStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryItemStatus struct {
	value *InventoryItemStatus
	isSet bool
}

func (v NullableInventoryItemStatus) Get() *InventoryItemStatus {
	return v.value
}

func (v *NullableInventoryItemStatus) Set(val *InventoryItemStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryItemStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryItemStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryItemStatus(val *InventoryItemStatus) *NullableInventoryItemStatus {
	return &NullableInventoryItemStatus{value: val, isSet: true}
}

func (v NullableInventoryItemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryItemStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}