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

// checks if the WritableCircuitRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableCircuitRequest{}

// WritableCircuitRequest Adds support for custom fields and tags.
type WritableCircuitRequest struct {
	// Unique circuit ID
	Cid             string                              `json:"cid"`
	Provider        BriefProviderRequest                `json:"provider"`
	ProviderAccount NullableBriefProviderAccountRequest `json:"provider_account,omitempty"`
	Type            BriefCircuitTypeRequest             `json:"type"`
	Status          *CircuitStatusValue                 `json:"status,omitempty"`
	Tenant          NullableBriefTenantRequest          `json:"tenant,omitempty"`
	InstallDate     NullableString                      `json:"install_date,omitempty"`
	TerminationDate NullableString                      `json:"termination_date,omitempty"`
	// Committed rate
	CommitRate           NullableInt32                                  `json:"commit_rate,omitempty"`
	Description          *string                                        `json:"description,omitempty"`
	Distance             NullableFloat64                                `json:"distance,omitempty"`
	DistanceUnit         NullableCircuitRequestDistanceUnit             `json:"distance_unit,omitempty"`
	Comments             *string                                        `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                             `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                         `json:"custom_fields,omitempty"`
	Assignments          []BriefCircuitGroupAssignmentSerializerRequest `json:"assignments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableCircuitRequest WritableCircuitRequest

// NewWritableCircuitRequest instantiates a new WritableCircuitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableCircuitRequest(cid string, provider BriefProviderRequest, type_ BriefCircuitTypeRequest) *WritableCircuitRequest {
	this := WritableCircuitRequest{}
	this.Cid = cid
	this.Provider = provider
	this.Type = type_
	return &this
}

// NewWritableCircuitRequestWithDefaults instantiates a new WritableCircuitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableCircuitRequestWithDefaults() *WritableCircuitRequest {
	this := WritableCircuitRequest{}
	return &this
}

// GetCid returns the Cid field value
func (o *WritableCircuitRequest) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *WritableCircuitRequest) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *WritableCircuitRequest) SetCid(v string) {
	o.Cid = v
}

// GetProvider returns the Provider field value
func (o *WritableCircuitRequest) GetProvider() BriefProviderRequest {
	if o == nil {
		var ret BriefProviderRequest
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *WritableCircuitRequest) GetProviderOk() (*BriefProviderRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *WritableCircuitRequest) SetProvider(v BriefProviderRequest) {
	o.Provider = v
}

// GetProviderAccount returns the ProviderAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitRequest) GetProviderAccount() BriefProviderAccountRequest {
	if o == nil || IsNil(o.ProviderAccount.Get()) {
		var ret BriefProviderAccountRequest
		return ret
	}
	return *o.ProviderAccount.Get()
}

// GetProviderAccountOk returns a tuple with the ProviderAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitRequest) GetProviderAccountOk() (*BriefProviderAccountRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderAccount.Get(), o.ProviderAccount.IsSet()
}

// HasProviderAccount returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasProviderAccount() bool {
	if o != nil && o.ProviderAccount.IsSet() {
		return true
	}

	return false
}

// SetProviderAccount gets a reference to the given NullableBriefProviderAccountRequest and assigns it to the ProviderAccount field.
func (o *WritableCircuitRequest) SetProviderAccount(v BriefProviderAccountRequest) {
	o.ProviderAccount.Set(&v)
}

// SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil
func (o *WritableCircuitRequest) SetProviderAccountNil() {
	o.ProviderAccount.Set(nil)
}

// UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
func (o *WritableCircuitRequest) UnsetProviderAccount() {
	o.ProviderAccount.Unset()
}

// GetType returns the Type field value
func (o *WritableCircuitRequest) GetType() BriefCircuitTypeRequest {
	if o == nil {
		var ret BriefCircuitTypeRequest
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WritableCircuitRequest) GetTypeOk() (*BriefCircuitTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WritableCircuitRequest) SetType(v BriefCircuitTypeRequest) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableCircuitRequest) GetStatus() CircuitStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret CircuitStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitRequest) GetStatusOk() (*CircuitStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CircuitStatusValue and assigns it to the Status field.
func (o *WritableCircuitRequest) SetStatus(v CircuitStatusValue) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *WritableCircuitRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableCircuitRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableCircuitRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetInstallDate returns the InstallDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitRequest) GetInstallDate() string {
	if o == nil || IsNil(o.InstallDate.Get()) {
		var ret string
		return ret
	}
	return *o.InstallDate.Get()
}

// GetInstallDateOk returns a tuple with the InstallDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitRequest) GetInstallDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallDate.Get(), o.InstallDate.IsSet()
}

// HasInstallDate returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasInstallDate() bool {
	if o != nil && o.InstallDate.IsSet() {
		return true
	}

	return false
}

// SetInstallDate gets a reference to the given NullableString and assigns it to the InstallDate field.
func (o *WritableCircuitRequest) SetInstallDate(v string) {
	o.InstallDate.Set(&v)
}

// SetInstallDateNil sets the value for InstallDate to be an explicit nil
func (o *WritableCircuitRequest) SetInstallDateNil() {
	o.InstallDate.Set(nil)
}

// UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
func (o *WritableCircuitRequest) UnsetInstallDate() {
	o.InstallDate.Unset()
}

// GetTerminationDate returns the TerminationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitRequest) GetTerminationDate() string {
	if o == nil || IsNil(o.TerminationDate.Get()) {
		var ret string
		return ret
	}
	return *o.TerminationDate.Get()
}

// GetTerminationDateOk returns a tuple with the TerminationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitRequest) GetTerminationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminationDate.Get(), o.TerminationDate.IsSet()
}

// HasTerminationDate returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasTerminationDate() bool {
	if o != nil && o.TerminationDate.IsSet() {
		return true
	}

	return false
}

// SetTerminationDate gets a reference to the given NullableString and assigns it to the TerminationDate field.
func (o *WritableCircuitRequest) SetTerminationDate(v string) {
	o.TerminationDate.Set(&v)
}

// SetTerminationDateNil sets the value for TerminationDate to be an explicit nil
func (o *WritableCircuitRequest) SetTerminationDateNil() {
	o.TerminationDate.Set(nil)
}

// UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
func (o *WritableCircuitRequest) UnsetTerminationDate() {
	o.TerminationDate.Unset()
}

// GetCommitRate returns the CommitRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitRequest) GetCommitRate() int32 {
	if o == nil || IsNil(o.CommitRate.Get()) {
		var ret int32
		return ret
	}
	return *o.CommitRate.Get()
}

// GetCommitRateOk returns a tuple with the CommitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitRequest) GetCommitRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitRate.Get(), o.CommitRate.IsSet()
}

// HasCommitRate returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasCommitRate() bool {
	if o != nil && o.CommitRate.IsSet() {
		return true
	}

	return false
}

// SetCommitRate gets a reference to the given NullableInt32 and assigns it to the CommitRate field.
func (o *WritableCircuitRequest) SetCommitRate(v int32) {
	o.CommitRate.Set(&v)
}

// SetCommitRateNil sets the value for CommitRate to be an explicit nil
func (o *WritableCircuitRequest) SetCommitRateNil() {
	o.CommitRate.Set(nil)
}

// UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
func (o *WritableCircuitRequest) UnsetCommitRate() {
	o.CommitRate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableCircuitRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableCircuitRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitRequest) GetDistance() float64 {
	if o == nil || IsNil(o.Distance.Get()) {
		var ret float64
		return ret
	}
	return *o.Distance.Get()
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitRequest) GetDistanceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Distance.Get(), o.Distance.IsSet()
}

// HasDistance returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasDistance() bool {
	if o != nil && o.Distance.IsSet() {
		return true
	}

	return false
}

// SetDistance gets a reference to the given NullableFloat64 and assigns it to the Distance field.
func (o *WritableCircuitRequest) SetDistance(v float64) {
	o.Distance.Set(&v)
}

// SetDistanceNil sets the value for Distance to be an explicit nil
func (o *WritableCircuitRequest) SetDistanceNil() {
	o.Distance.Set(nil)
}

// UnsetDistance ensures that no value is present for Distance, not even an explicit nil
func (o *WritableCircuitRequest) UnsetDistance() {
	o.Distance.Unset()
}

// GetDistanceUnit returns the DistanceUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableCircuitRequest) GetDistanceUnit() CircuitRequestDistanceUnit {
	if o == nil || IsNil(o.DistanceUnit.Get()) {
		var ret CircuitRequestDistanceUnit
		return ret
	}
	return *o.DistanceUnit.Get()
}

// GetDistanceUnitOk returns a tuple with the DistanceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableCircuitRequest) GetDistanceUnitOk() (*CircuitRequestDistanceUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistanceUnit.Get(), o.DistanceUnit.IsSet()
}

// HasDistanceUnit returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasDistanceUnit() bool {
	if o != nil && o.DistanceUnit.IsSet() {
		return true
	}

	return false
}

// SetDistanceUnit gets a reference to the given NullableCircuitRequestDistanceUnit and assigns it to the DistanceUnit field.
func (o *WritableCircuitRequest) SetDistanceUnit(v CircuitRequestDistanceUnit) {
	o.DistanceUnit.Set(&v)
}

// SetDistanceUnitNil sets the value for DistanceUnit to be an explicit nil
func (o *WritableCircuitRequest) SetDistanceUnitNil() {
	o.DistanceUnit.Set(nil)
}

// UnsetDistanceUnit ensures that no value is present for DistanceUnit, not even an explicit nil
func (o *WritableCircuitRequest) UnsetDistanceUnit() {
	o.DistanceUnit.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableCircuitRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableCircuitRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableCircuitRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableCircuitRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableCircuitRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableCircuitRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *WritableCircuitRequest) GetAssignments() []BriefCircuitGroupAssignmentSerializerRequest {
	if o == nil || IsNil(o.Assignments) {
		var ret []BriefCircuitGroupAssignmentSerializerRequest
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitRequest) GetAssignmentsOk() ([]BriefCircuitGroupAssignmentSerializerRequest, bool) {
	if o == nil || IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *WritableCircuitRequest) HasAssignments() bool {
	if o != nil && !IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []BriefCircuitGroupAssignmentSerializerRequest and assigns it to the Assignments field.
func (o *WritableCircuitRequest) SetAssignments(v []BriefCircuitGroupAssignmentSerializerRequest) {
	o.Assignments = v
}

func (o WritableCircuitRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableCircuitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["provider"] = o.Provider
	if o.ProviderAccount.IsSet() {
		toSerialize["provider_account"] = o.ProviderAccount.Get()
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.InstallDate.IsSet() {
		toSerialize["install_date"] = o.InstallDate.Get()
	}
	if o.TerminationDate.IsSet() {
		toSerialize["termination_date"] = o.TerminationDate.Get()
	}
	if o.CommitRate.IsSet() {
		toSerialize["commit_rate"] = o.CommitRate.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Distance.IsSet() {
		toSerialize["distance"] = o.Distance.Get()
	}
	if o.DistanceUnit.IsSet() {
		toSerialize["distance_unit"] = o.DistanceUnit.Get()
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
	if !IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableCircuitRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cid",
		"provider",
		"type",
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

	varWritableCircuitRequest := _WritableCircuitRequest{}

	err = json.Unmarshal(data, &varWritableCircuitRequest)

	if err != nil {
		return err
	}

	*o = WritableCircuitRequest(varWritableCircuitRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cid")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "provider_account")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "install_date")
		delete(additionalProperties, "termination_date")
		delete(additionalProperties, "commit_rate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "distance")
		delete(additionalProperties, "distance_unit")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "assignments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableCircuitRequest struct {
	value *WritableCircuitRequest
	isSet bool
}

func (v NullableWritableCircuitRequest) Get() *WritableCircuitRequest {
	return v.value
}

func (v *NullableWritableCircuitRequest) Set(val *WritableCircuitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableCircuitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableCircuitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableCircuitRequest(val *WritableCircuitRequest) *NullableWritableCircuitRequest {
	return &NullableWritableCircuitRequest{value: val, isSet: true}
}

func (v NullableWritableCircuitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableCircuitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
