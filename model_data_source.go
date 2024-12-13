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

// checks if the DataSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSource{}

// DataSource Adds support for custom fields and tags.
type DataSource struct {
	Id          int64            `json:"id"`
	Url         string           `json:"url"`
	DisplayUrl  string           `json:"display_url"`
	Display     string           `json:"display"`
	Name        string           `json:"name"`
	Type        DataSourceType   `json:"type"`
	SourceUrl   string           `json:"source_url"`
	Enabled     *bool            `json:"enabled,omitempty"`
	Status      DataSourceStatus `json:"status"`
	Description *string          `json:"description,omitempty"`
	Parameters  interface{}      `json:"parameters,omitempty"`
	// Patterns (one per line) matching files to ignore when syncing
	IgnoreRules          *string                `json:"ignore_rules,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	LastSynced           NullableTime           `json:"last_synced"`
	FileCount            int64                  `json:"file_count"`
	AdditionalProperties map[string]interface{}
}

type _DataSource DataSource

// NewDataSource instantiates a new DataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSource(id int64, url string, displayUrl string, display string, name string, type_ DataSourceType, sourceUrl string, status DataSourceStatus, created NullableTime, lastUpdated NullableTime, lastSynced NullableTime, fileCount int64) *DataSource {
	this := DataSource{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Type = type_
	this.SourceUrl = sourceUrl
	this.Status = status
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastSynced = lastSynced
	this.FileCount = fileCount
	return &this
}

// NewDataSourceWithDefaults instantiates a new DataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceWithDefaults() *DataSource {
	this := DataSource{}
	return &this
}

// GetId returns the Id field value
func (o *DataSource) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DataSource) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *DataSource) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DataSource) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *DataSource) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *DataSource) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *DataSource) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DataSource) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *DataSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataSource) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *DataSource) GetType() DataSourceType {
	if o == nil {
		var ret DataSourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetTypeOk() (*DataSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DataSource) SetType(v DataSourceType) {
	o.Type = v
}

// GetSourceUrl returns the SourceUrl field value
func (o *DataSource) GetSourceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetSourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceUrl, true
}

// SetSourceUrl sets field value
func (o *DataSource) SetSourceUrl(v string) {
	o.SourceUrl = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DataSource) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSource) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DataSource) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DataSource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatus returns the Status field value
func (o *DataSource) GetStatus() DataSourceStatus {
	if o == nil {
		var ret DataSourceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetStatusOk() (*DataSourceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DataSource) SetStatus(v DataSourceStatus) {
	o.Status = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DataSource) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSource) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DataSource) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DataSource) SetDescription(v string) {
	o.Description = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSource) GetParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSource) GetParametersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *DataSource) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given interface{} and assigns it to the Parameters field.
func (o *DataSource) SetParameters(v interface{}) {
	o.Parameters = v
}

// GetIgnoreRules returns the IgnoreRules field value if set, zero value otherwise.
func (o *DataSource) GetIgnoreRules() string {
	if o == nil || IsNil(o.IgnoreRules) {
		var ret string
		return ret
	}
	return *o.IgnoreRules
}

// GetIgnoreRulesOk returns a tuple with the IgnoreRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSource) GetIgnoreRulesOk() (*string, bool) {
	if o == nil || IsNil(o.IgnoreRules) {
		return nil, false
	}
	return o.IgnoreRules, true
}

// HasIgnoreRules returns a boolean if a field has been set.
func (o *DataSource) HasIgnoreRules() bool {
	if o != nil && !IsNil(o.IgnoreRules) {
		return true
	}

	return false
}

// SetIgnoreRules gets a reference to the given string and assigns it to the IgnoreRules field.
func (o *DataSource) SetIgnoreRules(v string) {
	o.IgnoreRules = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *DataSource) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSource) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *DataSource) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *DataSource) SetComments(v string) {
	o.Comments = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DataSource) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSource) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DataSource) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DataSource) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DataSource) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSource) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *DataSource) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DataSource) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSource) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *DataSource) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetLastSynced returns the LastSynced field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DataSource) GetLastSynced() time.Time {
	if o == nil || o.LastSynced.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastSynced.Get()
}

// GetLastSyncedOk returns a tuple with the LastSynced field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSource) GetLastSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSynced.Get(), o.LastSynced.IsSet()
}

// SetLastSynced sets field value
func (o *DataSource) SetLastSynced(v time.Time) {
	o.LastSynced.Set(&v)
}

// GetFileCount returns the FileCount field value
func (o *DataSource) GetFileCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileCount
}

// GetFileCountOk returns a tuple with the FileCount field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetFileCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileCount, true
}

// SetFileCount sets field value
func (o *DataSource) SetFileCount(v int64) {
	o.FileCount = v
}

func (o DataSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["source_url"] = o.SourceUrl
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.IgnoreRules) {
		toSerialize["ignore_rules"] = o.IgnoreRules
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["last_synced"] = o.LastSynced.Get()
	toSerialize["file_count"] = o.FileCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataSource) UnmarshalJSON(data []byte) (err error) {
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
		"source_url",
		"status",
		"created",
		"last_updated",
		"last_synced",
		"file_count",
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

	varDataSource := _DataSource{}

	err = json.Unmarshal(data, &varDataSource)

	if err != nil {
		return err
	}

	*o = DataSource(varDataSource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "source_url")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "status")
		delete(additionalProperties, "description")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "ignore_rules")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "last_synced")
		delete(additionalProperties, "file_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataSource struct {
	value *DataSource
	isSet bool
}

func (v NullableDataSource) Get() *DataSource {
	return v.value
}

func (v *NullableDataSource) Set(val *DataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSource(val *DataSource) *NullableDataSource {
	return &NullableDataSource{value: val, isSet: true}
}

func (v NullableDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
