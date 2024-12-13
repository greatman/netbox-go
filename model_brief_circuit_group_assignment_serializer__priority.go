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

// checks if the BriefCircuitGroupAssignmentSerializerPriority type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefCircuitGroupAssignmentSerializerPriority{}

// BriefCircuitGroupAssignmentSerializerPriority struct for BriefCircuitGroupAssignmentSerializerPriority
type BriefCircuitGroupAssignmentSerializerPriority struct {
	Value                *BriefCircuitGroupAssignmentSerializerPriorityValue `json:"value,omitempty"`
	Label                *BriefCircuitGroupAssignmentSerializerPriorityLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefCircuitGroupAssignmentSerializerPriority BriefCircuitGroupAssignmentSerializerPriority

// NewBriefCircuitGroupAssignmentSerializerPriority instantiates a new BriefCircuitGroupAssignmentSerializerPriority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefCircuitGroupAssignmentSerializerPriority() *BriefCircuitGroupAssignmentSerializerPriority {
	this := BriefCircuitGroupAssignmentSerializerPriority{}
	return &this
}

// NewBriefCircuitGroupAssignmentSerializerPriorityWithDefaults instantiates a new BriefCircuitGroupAssignmentSerializerPriority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefCircuitGroupAssignmentSerializerPriorityWithDefaults() *BriefCircuitGroupAssignmentSerializerPriority {
	this := BriefCircuitGroupAssignmentSerializerPriority{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BriefCircuitGroupAssignmentSerializerPriority) GetValue() BriefCircuitGroupAssignmentSerializerPriorityValue {
	if o == nil || IsNil(o.Value) {
		var ret BriefCircuitGroupAssignmentSerializerPriorityValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefCircuitGroupAssignmentSerializerPriority) GetValueOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BriefCircuitGroupAssignmentSerializerPriority) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given BriefCircuitGroupAssignmentSerializerPriorityValue and assigns it to the Value field.
func (o *BriefCircuitGroupAssignmentSerializerPriority) SetValue(v BriefCircuitGroupAssignmentSerializerPriorityValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BriefCircuitGroupAssignmentSerializerPriority) GetLabel() BriefCircuitGroupAssignmentSerializerPriorityLabel {
	if o == nil || IsNil(o.Label) {
		var ret BriefCircuitGroupAssignmentSerializerPriorityLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefCircuitGroupAssignmentSerializerPriority) GetLabelOk() (*BriefCircuitGroupAssignmentSerializerPriorityLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BriefCircuitGroupAssignmentSerializerPriority) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given BriefCircuitGroupAssignmentSerializerPriorityLabel and assigns it to the Label field.
func (o *BriefCircuitGroupAssignmentSerializerPriority) SetLabel(v BriefCircuitGroupAssignmentSerializerPriorityLabel) {
	o.Label = &v
}

func (o BriefCircuitGroupAssignmentSerializerPriority) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefCircuitGroupAssignmentSerializerPriority) ToMap() (map[string]interface{}, error) {
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

func (o *BriefCircuitGroupAssignmentSerializerPriority) UnmarshalJSON(data []byte) (err error) {
	varBriefCircuitGroupAssignmentSerializerPriority := _BriefCircuitGroupAssignmentSerializerPriority{}

	err = json.Unmarshal(data, &varBriefCircuitGroupAssignmentSerializerPriority)

	if err != nil {
		return err
	}

	*o = BriefCircuitGroupAssignmentSerializerPriority(varBriefCircuitGroupAssignmentSerializerPriority)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefCircuitGroupAssignmentSerializerPriority struct {
	value *BriefCircuitGroupAssignmentSerializerPriority
	isSet bool
}

func (v NullableBriefCircuitGroupAssignmentSerializerPriority) Get() *BriefCircuitGroupAssignmentSerializerPriority {
	return v.value
}

func (v *NullableBriefCircuitGroupAssignmentSerializerPriority) Set(val *BriefCircuitGroupAssignmentSerializerPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefCircuitGroupAssignmentSerializerPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefCircuitGroupAssignmentSerializerPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefCircuitGroupAssignmentSerializerPriority(val *BriefCircuitGroupAssignmentSerializerPriority) *NullableBriefCircuitGroupAssignmentSerializerPriority {
	return &NullableBriefCircuitGroupAssignmentSerializerPriority{value: val, isSet: true}
}

func (v NullableBriefCircuitGroupAssignmentSerializerPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefCircuitGroupAssignmentSerializerPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
