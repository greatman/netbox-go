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
	"time"
)

// checks if the VirtualCircuit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualCircuit{}

// VirtualCircuit Adds support for custom fields and tags.
type VirtualCircuit struct {
	Id         int32  `json:"id"`
	Url        string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display    string `json:"display"`
	// Unique circuit ID
	Cid                  string                       `json:"cid"`
	ProviderNetwork      BriefProviderNetwork         `json:"provider_network"`
	ProviderAccount      NullableBriefProviderAccount `json:"provider_account,omitempty"`
	Status               *CircuitStatus               `json:"status,omitempty"`
	Tenant               NullableBriefTenant          `json:"tenant,omitempty"`
	Description          *string                      `json:"description,omitempty"`
	Comments             *string                      `json:"comments,omitempty"`
	Tags                 []NestedTag                  `json:"tags,omitempty"`
	CustomFields         map[string]interface{}       `json:"custom_fields,omitempty"`
	Created              NullableTime                 `json:"created"`
	LastUpdated          NullableTime                 `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _VirtualCircuit VirtualCircuit

// NewVirtualCircuit instantiates a new VirtualCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualCircuit(id int32, url string, displayUrl string, display string, cid string, providerNetwork BriefProviderNetwork, created NullableTime, lastUpdated NullableTime) *VirtualCircuit {
	this := VirtualCircuit{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Cid = cid
	this.ProviderNetwork = providerNetwork
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewVirtualCircuitWithDefaults instantiates a new VirtualCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualCircuitWithDefaults() *VirtualCircuit {
	this := VirtualCircuit{}
	return &this
}

// GetId returns the Id field value
func (o *VirtualCircuit) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VirtualCircuit) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VirtualCircuit) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VirtualCircuit) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *VirtualCircuit) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *VirtualCircuit) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *VirtualCircuit) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VirtualCircuit) SetDisplay(v string) {
	o.Display = v
}

// GetCid returns the Cid field value
func (o *VirtualCircuit) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *VirtualCircuit) SetCid(v string) {
	o.Cid = v
}

// GetProviderNetwork returns the ProviderNetwork field value
func (o *VirtualCircuit) GetProviderNetwork() BriefProviderNetwork {
	if o == nil {
		var ret BriefProviderNetwork
		return ret
	}

	return o.ProviderNetwork
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetProviderNetworkOk() (*BriefProviderNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderNetwork, true
}

// SetProviderNetwork sets field value
func (o *VirtualCircuit) SetProviderNetwork(v BriefProviderNetwork) {
	o.ProviderNetwork = v
}

// GetProviderAccount returns the ProviderAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualCircuit) GetProviderAccount() BriefProviderAccount {
	if o == nil || IsNil(o.ProviderAccount.Get()) {
		var ret BriefProviderAccount
		return ret
	}
	return *o.ProviderAccount.Get()
}

// GetProviderAccountOk returns a tuple with the ProviderAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualCircuit) GetProviderAccountOk() (*BriefProviderAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderAccount.Get(), o.ProviderAccount.IsSet()
}

// HasProviderAccount returns a boolean if a field has been set.
func (o *VirtualCircuit) HasProviderAccount() bool {
	if o != nil && o.ProviderAccount.IsSet() {
		return true
	}

	return false
}

// SetProviderAccount gets a reference to the given NullableBriefProviderAccount and assigns it to the ProviderAccount field.
func (o *VirtualCircuit) SetProviderAccount(v BriefProviderAccount) {
	o.ProviderAccount.Set(&v)
}

// SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil
func (o *VirtualCircuit) SetProviderAccountNil() {
	o.ProviderAccount.Set(nil)
}

// UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
func (o *VirtualCircuit) UnsetProviderAccount() {
	o.ProviderAccount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualCircuit) GetStatus() CircuitStatus {
	if o == nil || IsNil(o.Status) {
		var ret CircuitStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetStatusOk() (*CircuitStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualCircuit) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CircuitStatus and assigns it to the Status field.
func (o *VirtualCircuit) SetStatus(v CircuitStatus) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualCircuit) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualCircuit) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *VirtualCircuit) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *VirtualCircuit) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *VirtualCircuit) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *VirtualCircuit) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualCircuit) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualCircuit) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualCircuit) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VirtualCircuit) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VirtualCircuit) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *VirtualCircuit) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualCircuit) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualCircuit) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *VirtualCircuit) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VirtualCircuit) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCircuit) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VirtualCircuit) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VirtualCircuit) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VirtualCircuit) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualCircuit) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *VirtualCircuit) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VirtualCircuit) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualCircuit) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *VirtualCircuit) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o VirtualCircuit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualCircuit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualCircuit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"cid",
		"provider_network",
		"created",
		"last_updated",
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

	varVirtualCircuit := _VirtualCircuit{}

	err = json.Unmarshal(data, &varVirtualCircuit)

	if err != nil {
		return err
	}

	*o = VirtualCircuit(varVirtualCircuit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "cid")
		delete(additionalProperties, "provider_network")
		delete(additionalProperties, "provider_account")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualCircuit struct {
	value *VirtualCircuit
	isSet bool
}

func (v NullableVirtualCircuit) Get() *VirtualCircuit {
	return v.value
}

func (v *NullableVirtualCircuit) Set(val *VirtualCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualCircuit(val *VirtualCircuit) *NullableVirtualCircuit {
	return &NullableVirtualCircuit{value: val, isSet: true}
}

func (v NullableVirtualCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
