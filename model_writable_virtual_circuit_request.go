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

// checks if the WritableVirtualCircuitRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableVirtualCircuitRequest{}

// WritableVirtualCircuitRequest Adds support for custom fields and tags.
type WritableVirtualCircuitRequest struct {
	// Unique circuit ID
	Cid                  string                              `json:"cid"`
	ProviderNetwork      BriefProviderNetworkRequest         `json:"provider_network"`
	ProviderAccount      NullableBriefProviderAccountRequest `json:"provider_account,omitempty"`
	Status               *CircuitStatusValue                 `json:"status,omitempty"`
	Tenant               NullableBriefTenantRequest          `json:"tenant,omitempty"`
	Description          *string                             `json:"description,omitempty"`
	Comments             *string                             `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                  `json:"tags,omitempty"`
	CustomFields         map[string]interface{}              `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableVirtualCircuitRequest WritableVirtualCircuitRequest

// NewWritableVirtualCircuitRequest instantiates a new WritableVirtualCircuitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVirtualCircuitRequest(cid string, providerNetwork BriefProviderNetworkRequest) *WritableVirtualCircuitRequest {
	this := WritableVirtualCircuitRequest{}
	this.Cid = cid
	this.ProviderNetwork = providerNetwork
	return &this
}

// NewWritableVirtualCircuitRequestWithDefaults instantiates a new WritableVirtualCircuitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVirtualCircuitRequestWithDefaults() *WritableVirtualCircuitRequest {
	this := WritableVirtualCircuitRequest{}
	return &this
}

// GetCid returns the Cid field value
func (o *WritableVirtualCircuitRequest) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualCircuitRequest) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *WritableVirtualCircuitRequest) SetCid(v string) {
	o.Cid = v
}

// GetProviderNetwork returns the ProviderNetwork field value
func (o *WritableVirtualCircuitRequest) GetProviderNetwork() BriefProviderNetworkRequest {
	if o == nil {
		var ret BriefProviderNetworkRequest
		return ret
	}

	return o.ProviderNetwork
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualCircuitRequest) GetProviderNetworkOk() (*BriefProviderNetworkRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderNetwork, true
}

// SetProviderNetwork sets field value
func (o *WritableVirtualCircuitRequest) SetProviderNetwork(v BriefProviderNetworkRequest) {
	o.ProviderNetwork = v
}

// GetProviderAccount returns the ProviderAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualCircuitRequest) GetProviderAccount() BriefProviderAccountRequest {
	if o == nil || IsNil(o.ProviderAccount.Get()) {
		var ret BriefProviderAccountRequest
		return ret
	}
	return *o.ProviderAccount.Get()
}

// GetProviderAccountOk returns a tuple with the ProviderAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualCircuitRequest) GetProviderAccountOk() (*BriefProviderAccountRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderAccount.Get(), o.ProviderAccount.IsSet()
}

// HasProviderAccount returns a boolean if a field has been set.
func (o *WritableVirtualCircuitRequest) HasProviderAccount() bool {
	if o != nil && o.ProviderAccount.IsSet() {
		return true
	}

	return false
}

// SetProviderAccount gets a reference to the given NullableBriefProviderAccountRequest and assigns it to the ProviderAccount field.
func (o *WritableVirtualCircuitRequest) SetProviderAccount(v BriefProviderAccountRequest) {
	o.ProviderAccount.Set(&v)
}

// SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil
func (o *WritableVirtualCircuitRequest) SetProviderAccountNil() {
	o.ProviderAccount.Set(nil)
}

// UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
func (o *WritableVirtualCircuitRequest) UnsetProviderAccount() {
	o.ProviderAccount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableVirtualCircuitRequest) GetStatus() CircuitStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret CircuitStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualCircuitRequest) GetStatusOk() (*CircuitStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableVirtualCircuitRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CircuitStatusValue and assigns it to the Status field.
func (o *WritableVirtualCircuitRequest) SetStatus(v CircuitStatusValue) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualCircuitRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualCircuitRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableVirtualCircuitRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *WritableVirtualCircuitRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableVirtualCircuitRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableVirtualCircuitRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableVirtualCircuitRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualCircuitRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableVirtualCircuitRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableVirtualCircuitRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableVirtualCircuitRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualCircuitRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableVirtualCircuitRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableVirtualCircuitRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVirtualCircuitRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualCircuitRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVirtualCircuitRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableVirtualCircuitRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableVirtualCircuitRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualCircuitRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableVirtualCircuitRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableVirtualCircuitRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableVirtualCircuitRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableVirtualCircuitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["provider_network"] = o.ProviderNetwork
	if o.ProviderAccount.IsSet() {
		toSerialize["provider_account"] = o.ProviderAccount.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *WritableVirtualCircuitRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cid",
		"provider_network",
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

	varWritableVirtualCircuitRequest := _WritableVirtualCircuitRequest{}

	err = json.Unmarshal(data, &varWritableVirtualCircuitRequest)

	if err != nil {
		return err
	}

	*o = WritableVirtualCircuitRequest(varWritableVirtualCircuitRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cid")
		delete(additionalProperties, "provider_network")
		delete(additionalProperties, "provider_account")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableVirtualCircuitRequest struct {
	value *WritableVirtualCircuitRequest
	isSet bool
}

func (v NullableWritableVirtualCircuitRequest) Get() *WritableVirtualCircuitRequest {
	return v.value
}

func (v *NullableWritableVirtualCircuitRequest) Set(val *WritableVirtualCircuitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVirtualCircuitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVirtualCircuitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVirtualCircuitRequest(val *WritableVirtualCircuitRequest) *NullableWritableVirtualCircuitRequest {
	return &NullableWritableVirtualCircuitRequest{value: val, isSet: true}
}

func (v NullableWritableVirtualCircuitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVirtualCircuitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}