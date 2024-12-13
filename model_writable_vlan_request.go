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

// checks if the WritableVLANRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableVLANRequest{}

// WritableVLANRequest Adds support for custom fields and tags.
type WritableVLANRequest struct {
	Site  NullableBriefSiteRequest      `json:"site,omitempty"`
	Group NullableBriefVLANGroupRequest `json:"group,omitempty"`
	// Numeric VLAN ID (1-4094)
	Vid                  int32                             `json:"vid"`
	Name                 string                            `json:"name"`
	Tenant               NullableBriefTenantRequest        `json:"tenant,omitempty"`
	Status               *PatchedWritableVLANRequestStatus `json:"status,omitempty"`
	Role                 NullableBriefRoleRequest          `json:"role,omitempty"`
	Description          *string                           `json:"description,omitempty"`
	QinqRole             NullableQInQRole                  `json:"qinq_role,omitempty"`
	QinqSvlan            NullableInt32                     `json:"qinq_svlan,omitempty"`
	Comments             *string                           `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                `json:"tags,omitempty"`
	CustomFields         map[string]interface{}            `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableVLANRequest WritableVLANRequest

// NewWritableVLANRequest instantiates a new WritableVLANRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVLANRequest(vid int32, name string) *WritableVLANRequest {
	this := WritableVLANRequest{}
	this.Vid = vid
	this.Name = name
	return &this
}

// NewWritableVLANRequestWithDefaults instantiates a new WritableVLANRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVLANRequestWithDefaults() *WritableVLANRequest {
	this := WritableVLANRequest{}
	return &this
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVLANRequest) GetSite() BriefSiteRequest {
	if o == nil || IsNil(o.Site.Get()) {
		var ret BriefSiteRequest
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVLANRequest) GetSiteOk() (*BriefSiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableBriefSiteRequest and assigns it to the Site field.
func (o *WritableVLANRequest) SetSite(v BriefSiteRequest) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *WritableVLANRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *WritableVLANRequest) UnsetSite() {
	o.Site.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVLANRequest) GetGroup() BriefVLANGroupRequest {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefVLANGroupRequest
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVLANRequest) GetGroupOk() (*BriefVLANGroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefVLANGroupRequest and assigns it to the Group field.
func (o *WritableVLANRequest) SetGroup(v BriefVLANGroupRequest) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *WritableVLANRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *WritableVLANRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetVid returns the Vid field value
func (o *WritableVLANRequest) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *WritableVLANRequest) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *WritableVLANRequest) SetVid(v int32) {
	o.Vid = v
}

// GetName returns the Name field value
func (o *WritableVLANRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableVLANRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableVLANRequest) SetName(v string) {
	o.Name = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVLANRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVLANRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *WritableVLANRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableVLANRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableVLANRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableVLANRequest) GetStatus() PatchedWritableVLANRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableVLANRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVLANRequest) GetStatusOk() (*PatchedWritableVLANRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableVLANRequestStatus and assigns it to the Status field.
func (o *WritableVLANRequest) SetStatus(v PatchedWritableVLANRequestStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVLANRequest) GetRole() BriefRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVLANRequest) GetRoleOk() (*BriefRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefRoleRequest and assigns it to the Role field.
func (o *WritableVLANRequest) SetRole(v BriefRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *WritableVLANRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *WritableVLANRequest) UnsetRole() {
	o.Role.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableVLANRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVLANRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableVLANRequest) SetDescription(v string) {
	o.Description = &v
}

// GetQinqRole returns the QinqRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVLANRequest) GetQinqRole() QInQRole {
	if o == nil || IsNil(o.QinqRole.Get()) {
		var ret QInQRole
		return ret
	}
	return *o.QinqRole.Get()
}

// GetQinqRoleOk returns a tuple with the QinqRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVLANRequest) GetQinqRoleOk() (*QInQRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.QinqRole.Get(), o.QinqRole.IsSet()
}

// HasQinqRole returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasQinqRole() bool {
	if o != nil && o.QinqRole.IsSet() {
		return true
	}

	return false
}

// SetQinqRole gets a reference to the given NullableQInQRole and assigns it to the QinqRole field.
func (o *WritableVLANRequest) SetQinqRole(v QInQRole) {
	o.QinqRole.Set(&v)
}

// SetQinqRoleNil sets the value for QinqRole to be an explicit nil
func (o *WritableVLANRequest) SetQinqRoleNil() {
	o.QinqRole.Set(nil)
}

// UnsetQinqRole ensures that no value is present for QinqRole, not even an explicit nil
func (o *WritableVLANRequest) UnsetQinqRole() {
	o.QinqRole.Unset()
}

// GetQinqSvlan returns the QinqSvlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVLANRequest) GetQinqSvlan() int32 {
	if o == nil || IsNil(o.QinqSvlan.Get()) {
		var ret int32
		return ret
	}
	return *o.QinqSvlan.Get()
}

// GetQinqSvlanOk returns a tuple with the QinqSvlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVLANRequest) GetQinqSvlanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QinqSvlan.Get(), o.QinqSvlan.IsSet()
}

// HasQinqSvlan returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasQinqSvlan() bool {
	if o != nil && o.QinqSvlan.IsSet() {
		return true
	}

	return false
}

// SetQinqSvlan gets a reference to the given NullableInt32 and assigns it to the QinqSvlan field.
func (o *WritableVLANRequest) SetQinqSvlan(v int32) {
	o.QinqSvlan.Set(&v)
}

// SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil
func (o *WritableVLANRequest) SetQinqSvlanNil() {
	o.QinqSvlan.Set(nil)
}

// UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
func (o *WritableVLANRequest) UnsetQinqSvlan() {
	o.QinqSvlan.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableVLANRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVLANRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableVLANRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVLANRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVLANRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableVLANRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableVLANRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVLANRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableVLANRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableVLANRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableVLANRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableVLANRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	toSerialize["vid"] = o.Vid
	toSerialize["name"] = o.Name
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.QinqRole.IsSet() {
		toSerialize["qinq_role"] = o.QinqRole.Get()
	}
	if o.QinqSvlan.IsSet() {
		toSerialize["qinq_svlan"] = o.QinqSvlan.Get()
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

func (o *WritableVLANRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vid",
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

	varWritableVLANRequest := _WritableVLANRequest{}

	err = json.Unmarshal(data, &varWritableVLANRequest)

	if err != nil {
		return err
	}

	*o = WritableVLANRequest(varWritableVLANRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "site")
		delete(additionalProperties, "group")
		delete(additionalProperties, "vid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "description")
		delete(additionalProperties, "qinq_role")
		delete(additionalProperties, "qinq_svlan")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableVLANRequest struct {
	value *WritableVLANRequest
	isSet bool
}

func (v NullableWritableVLANRequest) Get() *WritableVLANRequest {
	return v.value
}

func (v *NullableWritableVLANRequest) Set(val *WritableVLANRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVLANRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVLANRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVLANRequest(val *WritableVLANRequest) *NullableWritableVLANRequest {
	return &NullableWritableVLANRequest{value: val, isSet: true}
}

func (v NullableWritableVLANRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVLANRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
