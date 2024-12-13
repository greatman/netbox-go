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

// checks if the GroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRequest{}

// GroupRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type GroupRequest struct {
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	Permissions          []int64 `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupRequest GroupRequest

// NewGroupRequest instantiates a new GroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRequest(name string) *GroupRequest {
	this := GroupRequest{}
	this.Name = name
	return &this
}

// NewGroupRequestWithDefaults instantiates a new GroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRequestWithDefaults() *GroupRequest {
	this := GroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *GroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GroupRequest) GetPermissions() []int64 {
	if o == nil || IsNil(o.Permissions) {
		var ret []int64
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRequest) GetPermissionsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GroupRequest) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []int64 and assigns it to the Permissions field.
func (o *GroupRequest) SetPermissions(v []int64) {
	o.Permissions = v
}

func (o GroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varGroupRequest := _GroupRequest{}

	err = json.Unmarshal(data, &varGroupRequest)

	if err != nil {
		return err
	}

	*o = GroupRequest(varGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupRequest struct {
	value *GroupRequest
	isSet bool
}

func (v NullableGroupRequest) Get() *GroupRequest {
	return v.value
}

func (v *NullableGroupRequest) Set(val *GroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRequest(val *GroupRequest) *NullableGroupRequest {
	return &NullableGroupRequest{value: val, isSet: true}
}

func (v NullableGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
