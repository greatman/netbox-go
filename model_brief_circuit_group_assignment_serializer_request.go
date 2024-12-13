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

// checks if the BriefCircuitGroupAssignmentSerializerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefCircuitGroupAssignmentSerializerRequest{}

// BriefCircuitGroupAssignmentSerializerRequest Base serializer for group assignments under CircuitSerializer.
type BriefCircuitGroupAssignmentSerializerRequest struct {
	Group                BriefCircuitGroupRequest                            `json:"group"`
	Priority             *BriefCircuitGroupAssignmentSerializerPriorityValue `json:"priority,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefCircuitGroupAssignmentSerializerRequest BriefCircuitGroupAssignmentSerializerRequest

// NewBriefCircuitGroupAssignmentSerializerRequest instantiates a new BriefCircuitGroupAssignmentSerializerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefCircuitGroupAssignmentSerializerRequest(group BriefCircuitGroupRequest) *BriefCircuitGroupAssignmentSerializerRequest {
	this := BriefCircuitGroupAssignmentSerializerRequest{}
	this.Group = group
	return &this
}

// NewBriefCircuitGroupAssignmentSerializerRequestWithDefaults instantiates a new BriefCircuitGroupAssignmentSerializerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefCircuitGroupAssignmentSerializerRequestWithDefaults() *BriefCircuitGroupAssignmentSerializerRequest {
	this := BriefCircuitGroupAssignmentSerializerRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *BriefCircuitGroupAssignmentSerializerRequest) GetGroup() BriefCircuitGroupRequest {
	if o == nil {
		var ret BriefCircuitGroupRequest
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *BriefCircuitGroupAssignmentSerializerRequest) GetGroupOk() (*BriefCircuitGroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *BriefCircuitGroupAssignmentSerializerRequest) SetGroup(v BriefCircuitGroupRequest) {
	o.Group = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *BriefCircuitGroupAssignmentSerializerRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue {
	if o == nil || IsNil(o.Priority) {
		var ret BriefCircuitGroupAssignmentSerializerPriorityValue
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefCircuitGroupAssignmentSerializerRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *BriefCircuitGroupAssignmentSerializerRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given BriefCircuitGroupAssignmentSerializerPriorityValue and assigns it to the Priority field.
func (o *BriefCircuitGroupAssignmentSerializerRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue) {
	o.Priority = &v
}

func (o BriefCircuitGroupAssignmentSerializerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefCircuitGroupAssignmentSerializerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefCircuitGroupAssignmentSerializerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBriefCircuitGroupAssignmentSerializerRequest := _BriefCircuitGroupAssignmentSerializerRequest{}

	err = json.Unmarshal(data, &varBriefCircuitGroupAssignmentSerializerRequest)

	if err != nil {
		return err
	}

	*o = BriefCircuitGroupAssignmentSerializerRequest(varBriefCircuitGroupAssignmentSerializerRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefCircuitGroupAssignmentSerializerRequest struct {
	value *BriefCircuitGroupAssignmentSerializerRequest
	isSet bool
}

func (v NullableBriefCircuitGroupAssignmentSerializerRequest) Get() *BriefCircuitGroupAssignmentSerializerRequest {
	return v.value
}

func (v *NullableBriefCircuitGroupAssignmentSerializerRequest) Set(val *BriefCircuitGroupAssignmentSerializerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefCircuitGroupAssignmentSerializerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefCircuitGroupAssignmentSerializerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefCircuitGroupAssignmentSerializerRequest(val *BriefCircuitGroupAssignmentSerializerRequest) *NullableBriefCircuitGroupAssignmentSerializerRequest {
	return &NullableBriefCircuitGroupAssignmentSerializerRequest{value: val, isSet: true}
}

func (v NullableBriefCircuitGroupAssignmentSerializerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefCircuitGroupAssignmentSerializerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
