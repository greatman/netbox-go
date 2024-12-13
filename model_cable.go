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

// checks if the Cable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cable{}

// Cable Adds support for custom fields and tags.
type Cable struct {
	Id                   int32                   `json:"id"`
	Url                  string                  `json:"url"`
	DisplayUrl           string                  `json:"display_url"`
	Display              string                  `json:"display"`
	Type                 NullableCableType       `json:"type,omitempty"`
	ATerminations        []GenericObject         `json:"a_terminations,omitempty"`
	BTerminations        []GenericObject         `json:"b_terminations,omitempty"`
	Status               *CableStatus            `json:"status,omitempty"`
	Tenant               NullableBriefTenant     `json:"tenant,omitempty"`
	Label                *string                 `json:"label,omitempty"`
	Color                *string                 `json:"color,omitempty" validate:"regexp=^[0-9a-f]{6}$"`
	Length               NullableFloat64         `json:"length,omitempty"`
	LengthUnit           NullableCableLengthUnit `json:"length_unit,omitempty"`
	Description          *string                 `json:"description,omitempty"`
	Comments             *string                 `json:"comments,omitempty"`
	Tags                 []NestedTag             `json:"tags,omitempty"`
	CustomFields         map[string]interface{}  `json:"custom_fields,omitempty"`
	Created              NullableTime            `json:"created"`
	LastUpdated          NullableTime            `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _Cable Cable

// NewCable instantiates a new Cable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCable(id int32, url string, displayUrl string, display string, created NullableTime, lastUpdated NullableTime) *Cable {
	this := Cable{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewCableWithDefaults instantiates a new Cable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCableWithDefaults() *Cable {
	this := Cable{}
	return &this
}

// GetId returns the Id field value
func (o *Cable) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Cable) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Cable) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Cable) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Cable) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Cable) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *Cable) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *Cable) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *Cable) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *Cable) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Cable) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Cable) SetDisplay(v string) {
	o.Display = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cable) GetType() CableType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret CableType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cable) GetTypeOk() (*CableType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Cable) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableCableType and assigns it to the Type field.
func (o *Cable) SetType(v CableType) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *Cable) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Cable) UnsetType() {
	o.Type.Unset()
}

// GetATerminations returns the ATerminations field value if set, zero value otherwise.
func (o *Cable) GetATerminations() []GenericObject {
	if o == nil || IsNil(o.ATerminations) {
		var ret []GenericObject
		return ret
	}
	return o.ATerminations
}

// GetATerminationsOk returns a tuple with the ATerminations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cable) GetATerminationsOk() ([]GenericObject, bool) {
	if o == nil || IsNil(o.ATerminations) {
		return nil, false
	}
	return o.ATerminations, true
}

// HasATerminations returns a boolean if a field has been set.
func (o *Cable) HasATerminations() bool {
	if o != nil && !IsNil(o.ATerminations) {
		return true
	}

	return false
}

// SetATerminations gets a reference to the given []GenericObject and assigns it to the ATerminations field.
func (o *Cable) SetATerminations(v []GenericObject) {
	o.ATerminations = v
}

// GetBTerminations returns the BTerminations field value if set, zero value otherwise.
func (o *Cable) GetBTerminations() []GenericObject {
	if o == nil || IsNil(o.BTerminations) {
		var ret []GenericObject
		return ret
	}
	return o.BTerminations
}

// GetBTerminationsOk returns a tuple with the BTerminations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cable) GetBTerminationsOk() ([]GenericObject, bool) {
	if o == nil || IsNil(o.BTerminations) {
		return nil, false
	}
	return o.BTerminations, true
}

// HasBTerminations returns a boolean if a field has been set.
func (o *Cable) HasBTerminations() bool {
	if o != nil && !IsNil(o.BTerminations) {
		return true
	}

	return false
}

// SetBTerminations gets a reference to the given []GenericObject and assigns it to the BTerminations field.
func (o *Cable) SetBTerminations(v []GenericObject) {
	o.BTerminations = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Cable) GetStatus() CableStatus {
	if o == nil || IsNil(o.Status) {
		var ret CableStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cable) GetStatusOk() (*CableStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Cable) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CableStatus and assigns it to the Status field.
func (o *Cable) SetStatus(v CableStatus) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cable) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cable) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Cable) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *Cable) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Cable) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Cable) UnsetTenant() {
	o.Tenant.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Cable) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cable) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Cable) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Cable) SetLabel(v string) {
	o.Label = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Cable) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cable) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Cable) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Cable) SetColor(v string) {
	o.Color = &v
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cable) GetLength() float64 {
	if o == nil || IsNil(o.Length.Get()) {
		var ret float64
		return ret
	}
	return *o.Length.Get()
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cable) GetLengthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Length.Get(), o.Length.IsSet()
}

// HasLength returns a boolean if a field has been set.
func (o *Cable) HasLength() bool {
	if o != nil && o.Length.IsSet() {
		return true
	}

	return false
}

// SetLength gets a reference to the given NullableFloat64 and assigns it to the Length field.
func (o *Cable) SetLength(v float64) {
	o.Length.Set(&v)
}

// SetLengthNil sets the value for Length to be an explicit nil
func (o *Cable) SetLengthNil() {
	o.Length.Set(nil)
}

// UnsetLength ensures that no value is present for Length, not even an explicit nil
func (o *Cable) UnsetLength() {
	o.Length.Unset()
}

// GetLengthUnit returns the LengthUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cable) GetLengthUnit() CableLengthUnit {
	if o == nil || IsNil(o.LengthUnit.Get()) {
		var ret CableLengthUnit
		return ret
	}
	return *o.LengthUnit.Get()
}

// GetLengthUnitOk returns a tuple with the LengthUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cable) GetLengthUnitOk() (*CableLengthUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.LengthUnit.Get(), o.LengthUnit.IsSet()
}

// HasLengthUnit returns a boolean if a field has been set.
func (o *Cable) HasLengthUnit() bool {
	if o != nil && o.LengthUnit.IsSet() {
		return true
	}

	return false
}

// SetLengthUnit gets a reference to the given NullableCableLengthUnit and assigns it to the LengthUnit field.
func (o *Cable) SetLengthUnit(v CableLengthUnit) {
	o.LengthUnit.Set(&v)
}

// SetLengthUnitNil sets the value for LengthUnit to be an explicit nil
func (o *Cable) SetLengthUnitNil() {
	o.LengthUnit.Set(nil)
}

// UnsetLengthUnit ensures that no value is present for LengthUnit, not even an explicit nil
func (o *Cable) UnsetLengthUnit() {
	o.LengthUnit.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Cable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Cable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Cable) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Cable) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cable) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Cable) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Cable) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Cable) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cable) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Cable) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Cable) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Cable) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cable) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Cable) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Cable) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Cable) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cable) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Cable) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Cable) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cable) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Cable) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o Cable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.ATerminations) {
		toSerialize["a_terminations"] = o.ATerminations
	}
	if !IsNil(o.BTerminations) {
		toSerialize["b_terminations"] = o.BTerminations
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if o.Length.IsSet() {
		toSerialize["length"] = o.Length.Get()
	}
	if o.LengthUnit.IsSet() {
		toSerialize["length_unit"] = o.LengthUnit.Get()
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

func (o *Cable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
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

	varCable := _Cable{}

	err = json.Unmarshal(data, &varCable)

	if err != nil {
		return err
	}

	*o = Cable(varCable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "type")
		delete(additionalProperties, "a_terminations")
		delete(additionalProperties, "b_terminations")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "label")
		delete(additionalProperties, "color")
		delete(additionalProperties, "length")
		delete(additionalProperties, "length_unit")
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

type NullableCable struct {
	value *Cable
	isSet bool
}

func (v NullableCable) Get() *Cable {
	return v.value
}

func (v *NullableCable) Set(val *Cable) {
	v.value = val
	v.isSet = true
}

func (v NullableCable) IsSet() bool {
	return v.isSet
}

func (v *NullableCable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCable(val *Cable) *NullableCable {
	return &NullableCable{value: val, isSet: true}
}

func (v NullableCable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
