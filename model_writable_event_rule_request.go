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

// checks if the WritableEventRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableEventRuleRequest{}

// WritableEventRuleRequest Adds support for custom fields and tags.
type WritableEventRuleRequest struct {
	ObjectTypes []string `json:"object_types"`
	Name        string   `json:"name"`
	Enabled     *bool    `json:"enabled,omitempty"`
	// The types of event which will trigger this rule.
	EventTypes []EventRuleEventTypesInner `json:"event_types"`
	// A set of conditions which determine whether the event will be generated.
	Conditions           interface{}               `json:"conditions,omitempty"`
	ActionType           *EventRuleActionTypeValue `json:"action_type,omitempty"`
	ActionObjectType     string                    `json:"action_object_type"`
	ActionObjectId       NullableInt64             `json:"action_object_id,omitempty"`
	Description          *string                   `json:"description,omitempty"`
	CustomFields         map[string]interface{}    `json:"custom_fields,omitempty"`
	Tags                 []NestedTagRequest        `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableEventRuleRequest WritableEventRuleRequest

// NewWritableEventRuleRequest instantiates a new WritableEventRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableEventRuleRequest(objectTypes []string, name string, eventTypes []EventRuleEventTypesInner, actionObjectType string) *WritableEventRuleRequest {
	this := WritableEventRuleRequest{}
	this.ObjectTypes = objectTypes
	this.Name = name
	this.EventTypes = eventTypes
	this.ActionObjectType = actionObjectType
	return &this
}

// NewWritableEventRuleRequestWithDefaults instantiates a new WritableEventRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableEventRuleRequestWithDefaults() *WritableEventRuleRequest {
	this := WritableEventRuleRequest{}
	return &this
}

// GetObjectTypes returns the ObjectTypes field value
func (o *WritableEventRuleRequest) GetObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *WritableEventRuleRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *WritableEventRuleRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetName returns the Name field value
func (o *WritableEventRuleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableEventRuleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableEventRuleRequest) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WritableEventRuleRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableEventRuleRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WritableEventRuleRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WritableEventRuleRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypes returns the EventTypes field value
func (o *WritableEventRuleRequest) GetEventTypes() []EventRuleEventTypesInner {
	if o == nil {
		var ret []EventRuleEventTypesInner
		return ret
	}

	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value
// and a boolean to check if the value has been set.
func (o *WritableEventRuleRequest) GetEventTypesOk() ([]EventRuleEventTypesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// SetEventTypes sets field value
func (o *WritableEventRuleRequest) SetEventTypes(v []EventRuleEventTypesInner) {
	o.EventTypes = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableEventRuleRequest) GetConditions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableEventRuleRequest) GetConditionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *WritableEventRuleRequest) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given interface{} and assigns it to the Conditions field.
func (o *WritableEventRuleRequest) SetConditions(v interface{}) {
	o.Conditions = v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *WritableEventRuleRequest) GetActionType() EventRuleActionTypeValue {
	if o == nil || IsNil(o.ActionType) {
		var ret EventRuleActionTypeValue
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableEventRuleRequest) GetActionTypeOk() (*EventRuleActionTypeValue, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *WritableEventRuleRequest) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given EventRuleActionTypeValue and assigns it to the ActionType field.
func (o *WritableEventRuleRequest) SetActionType(v EventRuleActionTypeValue) {
	o.ActionType = &v
}

// GetActionObjectType returns the ActionObjectType field value
func (o *WritableEventRuleRequest) GetActionObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionObjectType
}

// GetActionObjectTypeOk returns a tuple with the ActionObjectType field value
// and a boolean to check if the value has been set.
func (o *WritableEventRuleRequest) GetActionObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionObjectType, true
}

// SetActionObjectType sets field value
func (o *WritableEventRuleRequest) SetActionObjectType(v string) {
	o.ActionObjectType = v
}

// GetActionObjectId returns the ActionObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableEventRuleRequest) GetActionObjectId() int64 {
	if o == nil || IsNil(o.ActionObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ActionObjectId.Get()
}

// GetActionObjectIdOk returns a tuple with the ActionObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableEventRuleRequest) GetActionObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionObjectId.Get(), o.ActionObjectId.IsSet()
}

// HasActionObjectId returns a boolean if a field has been set.
func (o *WritableEventRuleRequest) HasActionObjectId() bool {
	if o != nil && o.ActionObjectId.IsSet() {
		return true
	}

	return false
}

// SetActionObjectId gets a reference to the given NullableInt64 and assigns it to the ActionObjectId field.
func (o *WritableEventRuleRequest) SetActionObjectId(v int64) {
	o.ActionObjectId.Set(&v)
}

// SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil
func (o *WritableEventRuleRequest) SetActionObjectIdNil() {
	o.ActionObjectId.Set(nil)
}

// UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
func (o *WritableEventRuleRequest) UnsetActionObjectId() {
	o.ActionObjectId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableEventRuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableEventRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableEventRuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableEventRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableEventRuleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableEventRuleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableEventRuleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableEventRuleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableEventRuleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableEventRuleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableEventRuleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableEventRuleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o WritableEventRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableEventRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_types"] = o.ObjectTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["event_types"] = o.EventTypes
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.ActionType) {
		toSerialize["action_type"] = o.ActionType
	}
	toSerialize["action_object_type"] = o.ActionObjectType
	if o.ActionObjectId.IsSet() {
		toSerialize["action_object_id"] = o.ActionObjectId.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableEventRuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_types",
		"name",
		"event_types",
		"action_object_type",
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

	varWritableEventRuleRequest := _WritableEventRuleRequest{}

	err = json.Unmarshal(data, &varWritableEventRuleRequest)

	if err != nil {
		return err
	}

	*o = WritableEventRuleRequest(varWritableEventRuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "event_types")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "action_type")
		delete(additionalProperties, "action_object_type")
		delete(additionalProperties, "action_object_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableEventRuleRequest struct {
	value *WritableEventRuleRequest
	isSet bool
}

func (v NullableWritableEventRuleRequest) Get() *WritableEventRuleRequest {
	return v.value
}

func (v *NullableWritableEventRuleRequest) Set(val *WritableEventRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableEventRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableEventRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableEventRuleRequest(val *WritableEventRuleRequest) *NullableWritableEventRuleRequest {
	return &NullableWritableEventRuleRequest{value: val, isSet: true}
}

func (v NullableWritableEventRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableEventRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}