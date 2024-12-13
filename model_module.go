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

// checks if the Module type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Module{}

// Module Adds support for custom fields and tags.
type Module struct {
	Id         int64                `json:"id"`
	Url        string               `json:"url"`
	DisplayUrl string               `json:"display_url"`
	Display    string               `json:"display"`
	Device     BriefDevice          `json:"device"`
	ModuleBay  NestedModuleBay      `json:"module_bay"`
	ModuleType BriefModuleType      `json:"module_type"`
	Status     *InventoryItemStatus `json:"status,omitempty"`
	Serial     *string              `json:"serial,omitempty"`
	// A unique tag used to identify this device
	AssetTag             NullableString         `json:"asset_tag,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _Module Module

// NewModule instantiates a new Module object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModule(id int64, url string, displayUrl string, display string, device BriefDevice, moduleBay NestedModuleBay, moduleType BriefModuleType, created NullableTime, lastUpdated NullableTime) *Module {
	this := Module{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Device = device
	this.ModuleBay = moduleBay
	this.ModuleType = moduleType
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewModuleWithDefaults instantiates a new Module object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleWithDefaults() *Module {
	this := Module{}
	return &this
}

// GetId returns the Id field value
func (o *Module) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Module) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Module) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Module) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Module) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Module) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *Module) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *Module) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *Module) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *Module) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Module) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Module) SetDisplay(v string) {
	o.Display = v
}

// GetDevice returns the Device field value
func (o *Module) GetDevice() BriefDevice {
	if o == nil {
		var ret BriefDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *Module) GetDeviceOk() (*BriefDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *Module) SetDevice(v BriefDevice) {
	o.Device = v
}

// GetModuleBay returns the ModuleBay field value
func (o *Module) GetModuleBay() NestedModuleBay {
	if o == nil {
		var ret NestedModuleBay
		return ret
	}

	return o.ModuleBay
}

// GetModuleBayOk returns a tuple with the ModuleBay field value
// and a boolean to check if the value has been set.
func (o *Module) GetModuleBayOk() (*NestedModuleBay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleBay, true
}

// SetModuleBay sets field value
func (o *Module) SetModuleBay(v NestedModuleBay) {
	o.ModuleBay = v
}

// GetModuleType returns the ModuleType field value
func (o *Module) GetModuleType() BriefModuleType {
	if o == nil {
		var ret BriefModuleType
		return ret
	}

	return o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value
// and a boolean to check if the value has been set.
func (o *Module) GetModuleTypeOk() (*BriefModuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleType, true
}

// SetModuleType sets field value
func (o *Module) SetModuleType(v BriefModuleType) {
	o.ModuleType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Module) GetStatus() InventoryItemStatus {
	if o == nil || IsNil(o.Status) {
		var ret InventoryItemStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetStatusOk() (*InventoryItemStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Module) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given InventoryItemStatus and assigns it to the Status field.
func (o *Module) SetStatus(v InventoryItemStatus) {
	o.Status = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *Module) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *Module) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *Module) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *Module) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *Module) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *Module) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *Module) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Module) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Module) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Module) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Module) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Module) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Module) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Module) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Module) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Module) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Module) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Module) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Module) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Module) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Module) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Module) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Module) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o Module) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Module) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["device"] = o.Device
	toSerialize["module_bay"] = o.ModuleBay
	toSerialize["module_type"] = o.ModuleType
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
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

func (o *Module) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"device",
		"module_bay",
		"module_type",
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

	varModule := _Module{}

	err = json.Unmarshal(data, &varModule)

	if err != nil {
		return err
	}

	*o = Module(varModule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device")
		delete(additionalProperties, "module_bay")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
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

type NullableModule struct {
	value *Module
	isSet bool
}

func (v NullableModule) Get() *Module {
	return v.value
}

func (v *NullableModule) Set(val *Module) {
	v.value = val
	v.isSet = true
}

func (v NullableModule) IsSet() bool {
	return v.isSet
}

func (v *NullableModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModule(val *Module) *NullableModule {
	return &NullableModule{value: val, isSet: true}
}

func (v NullableModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
