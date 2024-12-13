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

// checks if the IKEPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IKEPolicyRequest{}

// IKEPolicyRequest Adds support for custom fields and tags.
type IKEPolicyRequest struct {
	Name                 string                 `json:"name"`
	Description          *string                `json:"description,omitempty"`
	Version              IKEPolicyVersionValue  `json:"version"`
	Mode                 *IKEPolicyModeValue    `json:"mode,omitempty"`
	Proposals            []int32                `json:"proposals,omitempty"`
	PresharedKey         *string                `json:"preshared_key,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IKEPolicyRequest IKEPolicyRequest

// NewIKEPolicyRequest instantiates a new IKEPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIKEPolicyRequest(name string, version IKEPolicyVersionValue) *IKEPolicyRequest {
	this := IKEPolicyRequest{}
	this.Name = name
	this.Version = version
	return &this
}

// NewIKEPolicyRequestWithDefaults instantiates a new IKEPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIKEPolicyRequestWithDefaults() *IKEPolicyRequest {
	this := IKEPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *IKEPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IKEPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IKEPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IKEPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IKEPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IKEPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value
func (o *IKEPolicyRequest) GetVersion() IKEPolicyVersionValue {
	if o == nil {
		var ret IKEPolicyVersionValue
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *IKEPolicyRequest) GetVersionOk() (*IKEPolicyVersionValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *IKEPolicyRequest) SetVersion(v IKEPolicyVersionValue) {
	o.Version = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *IKEPolicyRequest) GetMode() IKEPolicyModeValue {
	if o == nil || IsNil(o.Mode) {
		var ret IKEPolicyModeValue
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicyRequest) GetModeOk() (*IKEPolicyModeValue, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *IKEPolicyRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given IKEPolicyModeValue and assigns it to the Mode field.
func (o *IKEPolicyRequest) SetMode(v IKEPolicyModeValue) {
	o.Mode = &v
}

// GetProposals returns the Proposals field value if set, zero value otherwise.
func (o *IKEPolicyRequest) GetProposals() []int32 {
	if o == nil || IsNil(o.Proposals) {
		var ret []int32
		return ret
	}
	return o.Proposals
}

// GetProposalsOk returns a tuple with the Proposals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicyRequest) GetProposalsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Proposals) {
		return nil, false
	}
	return o.Proposals, true
}

// HasProposals returns a boolean if a field has been set.
func (o *IKEPolicyRequest) HasProposals() bool {
	if o != nil && !IsNil(o.Proposals) {
		return true
	}

	return false
}

// SetProposals gets a reference to the given []int32 and assigns it to the Proposals field.
func (o *IKEPolicyRequest) SetProposals(v []int32) {
	o.Proposals = v
}

// GetPresharedKey returns the PresharedKey field value if set, zero value otherwise.
func (o *IKEPolicyRequest) GetPresharedKey() string {
	if o == nil || IsNil(o.PresharedKey) {
		var ret string
		return ret
	}
	return *o.PresharedKey
}

// GetPresharedKeyOk returns a tuple with the PresharedKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicyRequest) GetPresharedKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PresharedKey) {
		return nil, false
	}
	return o.PresharedKey, true
}

// HasPresharedKey returns a boolean if a field has been set.
func (o *IKEPolicyRequest) HasPresharedKey() bool {
	if o != nil && !IsNil(o.PresharedKey) {
		return true
	}

	return false
}

// SetPresharedKey gets a reference to the given string and assigns it to the PresharedKey field.
func (o *IKEPolicyRequest) SetPresharedKey(v string) {
	o.PresharedKey = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IKEPolicyRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicyRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IKEPolicyRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IKEPolicyRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IKEPolicyRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicyRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IKEPolicyRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *IKEPolicyRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IKEPolicyRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicyRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IKEPolicyRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *IKEPolicyRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o IKEPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IKEPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Proposals) {
		toSerialize["proposals"] = o.Proposals
	}
	if !IsNil(o.PresharedKey) {
		toSerialize["preshared_key"] = o.PresharedKey
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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

func (o *IKEPolicyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"version",
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

	varIKEPolicyRequest := _IKEPolicyRequest{}

	err = json.Unmarshal(data, &varIKEPolicyRequest)

	if err != nil {
		return err
	}

	*o = IKEPolicyRequest(varIKEPolicyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "version")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "proposals")
		delete(additionalProperties, "preshared_key")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIKEPolicyRequest struct {
	value *IKEPolicyRequest
	isSet bool
}

func (v NullableIKEPolicyRequest) Get() *IKEPolicyRequest {
	return v.value
}

func (v *NullableIKEPolicyRequest) Set(val *IKEPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEPolicyRequest(val *IKEPolicyRequest) *NullableIKEPolicyRequest {
	return &NullableIKEPolicyRequest{value: val, isSet: true}
}

func (v NullableIKEPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
