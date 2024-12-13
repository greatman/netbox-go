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

// checks if the ObjectType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectType{}

// ObjectType struct for ObjectType
type ObjectType struct {
	Id                   int64  `json:"id"`
	Url                  string `json:"url"`
	Display              string `json:"display"`
	AppLabel             string `json:"app_label"`
	Model                string `json:"model"`
	AdditionalProperties map[string]interface{}
}

type _ObjectType ObjectType

// NewObjectType instantiates a new ObjectType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectType(id int64, url string, display string, appLabel string, model string) *ObjectType {
	this := ObjectType{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.AppLabel = appLabel
	this.Model = model
	return &this
}

// NewObjectTypeWithDefaults instantiates a new ObjectType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectTypeWithDefaults() *ObjectType {
	this := ObjectType{}
	return &this
}

// GetId returns the Id field value
func (o *ObjectType) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObjectType) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObjectType) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ObjectType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ObjectType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ObjectType) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ObjectType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ObjectType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ObjectType) SetDisplay(v string) {
	o.Display = v
}

// GetAppLabel returns the AppLabel field value
func (o *ObjectType) GetAppLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppLabel
}

// GetAppLabelOk returns a tuple with the AppLabel field value
// and a boolean to check if the value has been set.
func (o *ObjectType) GetAppLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLabel, true
}

// SetAppLabel sets field value
func (o *ObjectType) SetAppLabel(v string) {
	o.AppLabel = v
}

// GetModel returns the Model field value
func (o *ObjectType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ObjectType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ObjectType) SetModel(v string) {
	o.Model = v
}

func (o ObjectType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["app_label"] = o.AppLabel
	toSerialize["model"] = o.Model

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ObjectType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"app_label",
		"model",
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

	varObjectType := _ObjectType{}

	err = json.Unmarshal(data, &varObjectType)

	if err != nil {
		return err
	}

	*o = ObjectType(varObjectType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "app_label")
		delete(additionalProperties, "model")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjectType struct {
	value *ObjectType
	isSet bool
}

func (v NullableObjectType) Get() *ObjectType {
	return v.value
}

func (v *NullableObjectType) Set(val *ObjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectType(val *ObjectType) *NullableObjectType {
	return &NullableObjectType{value: val, isSet: true}
}

func (v NullableObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
