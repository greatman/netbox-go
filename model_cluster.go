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

// checks if the Cluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cluster{}

// Cluster Adds support for custom fields and tags.
type Cluster struct {
	Id                   int64                     `json:"id"`
	Url                  string                    `json:"url"`
	DisplayUrl           string                    `json:"display_url"`
	Display              string                    `json:"display"`
	Name                 string                    `json:"name"`
	Type                 BriefClusterType          `json:"type"`
	Group                NullableBriefClusterGroup `json:"group,omitempty"`
	Status               *ClusterStatus            `json:"status,omitempty"`
	Tenant               NullableBriefTenant       `json:"tenant,omitempty"`
	ScopeType            NullableString            `json:"scope_type,omitempty"`
	ScopeId              NullableInt64             `json:"scope_id,omitempty"`
	Scope                interface{}               `json:"scope"`
	Description          *string                   `json:"description,omitempty"`
	Comments             *string                   `json:"comments,omitempty"`
	Tags                 []NestedTag               `json:"tags,omitempty"`
	CustomFields         map[string]interface{}    `json:"custom_fields,omitempty"`
	Created              NullableTime              `json:"created"`
	LastUpdated          NullableTime              `json:"last_updated"`
	DeviceCount          int64                     `json:"device_count"`
	VirtualmachineCount  int64                     `json:"virtualmachine_count"`
	AllocatedVcpus       float64                   `json:"allocated_vcpus"`
	AllocatedMemory      int64                     `json:"allocated_memory"`
	AllocatedDisk        int64                     `json:"allocated_disk"`
	AdditionalProperties map[string]interface{}
}

type _Cluster Cluster

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(id int64, url string, displayUrl string, display string, name string, type_ BriefClusterType, scope interface{}, created NullableTime, lastUpdated NullableTime, deviceCount int64, virtualmachineCount int64, allocatedVcpus float64, allocatedMemory int64, allocatedDisk int64) *Cluster {
	this := Cluster{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Type = type_
	this.Scope = scope
	this.Created = created
	this.LastUpdated = lastUpdated
	this.DeviceCount = deviceCount
	this.VirtualmachineCount = virtualmachineCount
	this.AllocatedVcpus = allocatedVcpus
	this.AllocatedMemory = allocatedMemory
	this.AllocatedDisk = allocatedDisk
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	return &this
}

// GetId returns the Id field value
func (o *Cluster) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Cluster) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Cluster) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Cluster) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *Cluster) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *Cluster) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *Cluster) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Cluster) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Cluster) GetType() BriefClusterType {
	if o == nil {
		var ret BriefClusterType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetTypeOk() (*BriefClusterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Cluster) SetType(v BriefClusterType) {
	o.Type = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetGroup() BriefClusterGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefClusterGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetGroupOk() (*BriefClusterGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *Cluster) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefClusterGroup and assigns it to the Group field.
func (o *Cluster) SetGroup(v BriefClusterGroup) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *Cluster) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *Cluster) UnsetGroup() {
	o.Group.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Cluster) GetStatus() ClusterStatus {
	if o == nil || IsNil(o.Status) {
		var ret ClusterStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetStatusOk() (*ClusterStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Cluster) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterStatus and assigns it to the Status field.
func (o *Cluster) SetStatus(v ClusterStatus) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Cluster) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *Cluster) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Cluster) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Cluster) UnsetTenant() {
	o.Tenant.Unset()
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetScopeType() string {
	if o == nil || IsNil(o.ScopeType.Get()) {
		var ret string
		return ret
	}
	return *o.ScopeType.Get()
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetScopeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeType.Get(), o.ScopeType.IsSet()
}

// HasScopeType returns a boolean if a field has been set.
func (o *Cluster) HasScopeType() bool {
	if o != nil && o.ScopeType.IsSet() {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given NullableString and assigns it to the ScopeType field.
func (o *Cluster) SetScopeType(v string) {
	o.ScopeType.Set(&v)
}

// SetScopeTypeNil sets the value for ScopeType to be an explicit nil
func (o *Cluster) SetScopeTypeNil() {
	o.ScopeType.Set(nil)
}

// UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
func (o *Cluster) UnsetScopeType() {
	o.ScopeType.Unset()
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetScopeId() int64 {
	if o == nil || IsNil(o.ScopeId.Get()) {
		var ret int64
		return ret
	}
	return *o.ScopeId.Get()
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetScopeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeId.Get(), o.ScopeId.IsSet()
}

// HasScopeId returns a boolean if a field has been set.
func (o *Cluster) HasScopeId() bool {
	if o != nil && o.ScopeId.IsSet() {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given NullableInt64 and assigns it to the ScopeId field.
func (o *Cluster) SetScopeId(v int64) {
	o.ScopeId.Set(&v)
}

// SetScopeIdNil sets the value for ScopeId to be an explicit nil
func (o *Cluster) SetScopeIdNil() {
	o.ScopeId.Set(nil)
}

// UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
func (o *Cluster) UnsetScopeId() {
	o.ScopeId.Unset()
}

// GetScope returns the Scope field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Cluster) GetScope() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetScopeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *Cluster) SetScope(v interface{}) {
	o.Scope = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Cluster) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Cluster) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Cluster) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Cluster) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Cluster) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Cluster) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Cluster) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Cluster) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Cluster) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Cluster) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Cluster) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Cluster) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Cluster) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Cluster) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Cluster) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cluster) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Cluster) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetDeviceCount returns the DeviceCount field value
func (o *Cluster) GetDeviceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetDeviceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *Cluster) SetDeviceCount(v int64) {
	o.DeviceCount = v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value
func (o *Cluster) GetVirtualmachineCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetVirtualmachineCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualmachineCount, true
}

// SetVirtualmachineCount sets field value
func (o *Cluster) SetVirtualmachineCount(v int64) {
	o.VirtualmachineCount = v
}

// GetAllocatedVcpus returns the AllocatedVcpus field value
func (o *Cluster) GetAllocatedVcpus() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AllocatedVcpus
}

// GetAllocatedVcpusOk returns a tuple with the AllocatedVcpus field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetAllocatedVcpusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocatedVcpus, true
}

// SetAllocatedVcpus sets field value
func (o *Cluster) SetAllocatedVcpus(v float64) {
	o.AllocatedVcpus = v
}

// GetAllocatedMemory returns the AllocatedMemory field value
func (o *Cluster) GetAllocatedMemory() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AllocatedMemory
}

// GetAllocatedMemoryOk returns a tuple with the AllocatedMemory field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetAllocatedMemoryOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocatedMemory, true
}

// SetAllocatedMemory sets field value
func (o *Cluster) SetAllocatedMemory(v int64) {
	o.AllocatedMemory = v
}

// GetAllocatedDisk returns the AllocatedDisk field value
func (o *Cluster) GetAllocatedDisk() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AllocatedDisk
}

// GetAllocatedDiskOk returns a tuple with the AllocatedDisk field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetAllocatedDiskOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocatedDisk, true
}

// SetAllocatedDisk sets field value
func (o *Cluster) SetAllocatedDisk(v int64) {
	o.AllocatedDisk = v
}

func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.ScopeType.IsSet() {
		toSerialize["scope_type"] = o.ScopeType.Get()
	}
	if o.ScopeId.IsSet() {
		toSerialize["scope_id"] = o.ScopeId.Get()
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
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
	toSerialize["device_count"] = o.DeviceCount
	toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	toSerialize["allocated_vcpus"] = o.AllocatedVcpus
	toSerialize["allocated_memory"] = o.AllocatedMemory
	toSerialize["allocated_disk"] = o.AllocatedDisk

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Cluster) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"name",
		"type",
		"scope",
		"created",
		"last_updated",
		"device_count",
		"virtualmachine_count",
		"allocated_vcpus",
		"allocated_memory",
		"allocated_disk",
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

	varCluster := _Cluster{}

	err = json.Unmarshal(data, &varCluster)

	if err != nil {
		return err
	}

	*o = Cluster(varCluster)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "group")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "scope_type")
		delete(additionalProperties, "scope_id")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "virtualmachine_count")
		delete(additionalProperties, "allocated_vcpus")
		delete(additionalProperties, "allocated_memory")
		delete(additionalProperties, "allocated_disk")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
