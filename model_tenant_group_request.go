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

// checks if the TenantGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantGroupRequest{}

// TenantGroupRequest Extends PrimaryModelSerializer to include MPTT support.
type TenantGroupRequest struct {
	Name                 string                           `json:"name"`
	Slug                 string                           `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Parent               NullableNestedTenantGroupRequest `json:"parent,omitempty"`
	Description          *string                          `json:"description,omitempty"`
	Tags                 []NestedTagRequest               `json:"tags,omitempty"`
	CustomFields         map[string]interface{}           `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TenantGroupRequest TenantGroupRequest

// NewTenantGroupRequest instantiates a new TenantGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantGroupRequest(name string, slug string) *TenantGroupRequest {
	this := TenantGroupRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewTenantGroupRequestWithDefaults instantiates a new TenantGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantGroupRequestWithDefaults() *TenantGroupRequest {
	this := TenantGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *TenantGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TenantGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TenantGroupRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *TenantGroupRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *TenantGroupRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *TenantGroupRequest) SetSlug(v string) {
	o.Slug = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantGroupRequest) GetParent() NestedTenantGroupRequest {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret NestedTenantGroupRequest
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantGroupRequest) GetParentOk() (*NestedTenantGroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *TenantGroupRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableNestedTenantGroupRequest and assigns it to the Parent field.
func (o *TenantGroupRequest) SetParent(v NestedTenantGroupRequest) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *TenantGroupRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *TenantGroupRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TenantGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TenantGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TenantGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TenantGroupRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantGroupRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TenantGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *TenantGroupRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *TenantGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TenantGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *TenantGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o TenantGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TenantGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varTenantGroupRequest := _TenantGroupRequest{}

	err = json.Unmarshal(data, &varTenantGroupRequest)

	if err != nil {
		return err
	}

	*o = TenantGroupRequest(varTenantGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTenantGroupRequest struct {
	value *TenantGroupRequest
	isSet bool
}

func (v NullableTenantGroupRequest) Get() *TenantGroupRequest {
	return v.value
}

func (v *NullableTenantGroupRequest) Set(val *TenantGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantGroupRequest(val *TenantGroupRequest) *NullableTenantGroupRequest {
	return &NullableTenantGroupRequest{value: val, isSet: true}
}

func (v NullableTenantGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
