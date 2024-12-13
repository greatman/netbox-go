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

// checks if the ModuleBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleBayRequest{}

// ModuleBayRequest Adds support for custom fields and tags.
type ModuleBayRequest struct {
	Device          BriefDeviceRequest         `json:"device"`
	Module          NullableBriefModuleRequest `json:"module,omitempty"`
	Name            string                     `json:"name"`
	InstalledModule NullableBriefModuleRequest `json:"installed_module,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// Identifier to reference when renaming installed components
	Position             *string                `json:"position,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleBayRequest ModuleBayRequest

// NewModuleBayRequest instantiates a new ModuleBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleBayRequest(device BriefDeviceRequest, name string) *ModuleBayRequest {
	this := ModuleBayRequest{}
	this.Device = device
	this.Name = name
	return &this
}

// NewModuleBayRequestWithDefaults instantiates a new ModuleBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleBayRequestWithDefaults() *ModuleBayRequest {
	this := ModuleBayRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *ModuleBayRequest) GetDevice() BriefDeviceRequest {
	if o == nil {
		var ret BriefDeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *ModuleBayRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleBayRequest) GetModule() BriefModuleRequest {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BriefModuleRequest
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleBayRequest) GetModuleOk() (*BriefModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBriefModuleRequest and assigns it to the Module field.
func (o *ModuleBayRequest) SetModule(v BriefModuleRequest) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *ModuleBayRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *ModuleBayRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value
func (o *ModuleBayRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModuleBayRequest) SetName(v string) {
	o.Name = v
}

// GetInstalledModule returns the InstalledModule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleBayRequest) GetInstalledModule() BriefModuleRequest {
	if o == nil || IsNil(o.InstalledModule.Get()) {
		var ret BriefModuleRequest
		return ret
	}
	return *o.InstalledModule.Get()
}

// GetInstalledModuleOk returns a tuple with the InstalledModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleBayRequest) GetInstalledModuleOk() (*BriefModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledModule.Get(), o.InstalledModule.IsSet()
}

// HasInstalledModule returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasInstalledModule() bool {
	if o != nil && o.InstalledModule.IsSet() {
		return true
	}

	return false
}

// SetInstalledModule gets a reference to the given NullableBriefModuleRequest and assigns it to the InstalledModule field.
func (o *ModuleBayRequest) SetInstalledModule(v BriefModuleRequest) {
	o.InstalledModule.Set(&v)
}

// SetInstalledModuleNil sets the value for InstalledModule to be an explicit nil
func (o *ModuleBayRequest) SetInstalledModuleNil() {
	o.InstalledModule.Set(nil)
}

// UnsetInstalledModule ensures that no value is present for InstalledModule, not even an explicit nil
func (o *ModuleBayRequest) UnsetInstalledModule() {
	o.InstalledModule.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ModuleBayRequest) SetLabel(v string) {
	o.Label = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ModuleBayRequest) SetPosition(v string) {
	o.Position = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModuleBayRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ModuleBayRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ModuleBayRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ModuleBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["name"] = o.Name
	if o.InstalledModule.IsSet() {
		toSerialize["installed_module"] = o.InstalledModule.Get()
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
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

func (o *ModuleBayRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
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

	varModuleBayRequest := _ModuleBayRequest{}

	err = json.Unmarshal(data, &varModuleBayRequest)

	if err != nil {
		return err
	}

	*o = ModuleBayRequest(varModuleBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "installed_module")
		delete(additionalProperties, "label")
		delete(additionalProperties, "position")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleBayRequest struct {
	value *ModuleBayRequest
	isSet bool
}

func (v NullableModuleBayRequest) Get() *ModuleBayRequest {
	return v.value
}

func (v *NullableModuleBayRequest) Set(val *ModuleBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleBayRequest(val *ModuleBayRequest) *NullableModuleBayRequest {
	return &NullableModuleBayRequest{value: val, isSet: true}
}

func (v NullableModuleBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
