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

// checks if the Script type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Script{}

// Script Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type Script struct {
	Id                   int64          `json:"id"`
	Url                  string         `json:"url"`
	DisplayUrl           string         `json:"display_url"`
	Module               int64          `json:"module"`
	Name                 string         `json:"name"`
	Description          NullableString `json:"description"`
	Vars                 interface{}    `json:"vars"`
	Result               BriefJob       `json:"result"`
	Display              string         `json:"display"`
	IsExecutable         bool           `json:"is_executable"`
	AdditionalProperties map[string]interface{}
}

type _Script Script

// NewScript instantiates a new Script object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScript(id int64, url string, displayUrl string, module int64, name string, description NullableString, vars interface{}, result BriefJob, display string, isExecutable bool) *Script {
	this := Script{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Module = module
	this.Name = name
	this.Description = description
	this.Vars = vars
	this.Result = result
	this.Display = display
	this.IsExecutable = isExecutable
	return &this
}

// NewScriptWithDefaults instantiates a new Script object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptWithDefaults() *Script {
	this := Script{}
	return &this
}

// GetId returns the Id field value
func (o *Script) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Script) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Script) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Script) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Script) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Script) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *Script) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *Script) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *Script) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetModule returns the Module field value
func (o *Script) GetModule() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *Script) GetModuleOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *Script) SetModule(v int64) {
	o.Module = v
}

// GetName returns the Name field value
func (o *Script) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Script) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Script) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Script) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Script) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Script) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetVars returns the Vars field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Script) GetVars() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Vars
}

// GetVarsOk returns a tuple with the Vars field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Script) GetVarsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Vars) {
		return nil, false
	}
	return &o.Vars, true
}

// SetVars sets field value
func (o *Script) SetVars(v interface{}) {
	o.Vars = v
}

// GetResult returns the Result field value
func (o *Script) GetResult() BriefJob {
	if o == nil {
		var ret BriefJob
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *Script) GetResultOk() (*BriefJob, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *Script) SetResult(v BriefJob) {
	o.Result = v
}

// GetDisplay returns the Display field value
func (o *Script) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Script) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Script) SetDisplay(v string) {
	o.Display = v
}

// GetIsExecutable returns the IsExecutable field value
func (o *Script) GetIsExecutable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExecutable
}

// GetIsExecutableOk returns a tuple with the IsExecutable field value
// and a boolean to check if the value has been set.
func (o *Script) GetIsExecutableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExecutable, true
}

// SetIsExecutable sets field value
func (o *Script) SetIsExecutable(v bool) {
	o.IsExecutable = v
}

func (o Script) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Script) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["module"] = o.Module
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	if o.Vars != nil {
		toSerialize["vars"] = o.Vars
	}
	toSerialize["result"] = o.Result
	toSerialize["display"] = o.Display
	toSerialize["is_executable"] = o.IsExecutable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Script) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"module",
		"name",
		"description",
		"vars",
		"result",
		"display",
		"is_executable",
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

	varScript := _Script{}

	err = json.Unmarshal(data, &varScript)

	if err != nil {
		return err
	}

	*o = Script(varScript)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "vars")
		delete(additionalProperties, "result")
		delete(additionalProperties, "display")
		delete(additionalProperties, "is_executable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScript struct {
	value *Script
	isSet bool
}

func (v NullableScript) Get() *Script {
	return v.value
}

func (v *NullableScript) Set(val *Script) {
	v.value = val
	v.isSet = true
}

func (v NullableScript) IsSet() bool {
	return v.isSet
}

func (v *NullableScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScript(val *Script) *NullableScript {
	return &NullableScript{value: val, isSet: true}
}

func (v NullableScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
