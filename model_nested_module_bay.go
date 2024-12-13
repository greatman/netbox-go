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

// checks if the NestedModuleBay type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedModuleBay{}

// NestedModuleBay Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedModuleBay struct {
	Id                   int64  `json:"id"`
	Url                  string `json:"url"`
	DisplayUrl           string `json:"display_url"`
	Display              string `json:"display"`
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedModuleBay NestedModuleBay

// NewNestedModuleBay instantiates a new NestedModuleBay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedModuleBay(id int64, url string, displayUrl string, display string, name string) *NestedModuleBay {
	this := NestedModuleBay{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	return &this
}

// NewNestedModuleBayWithDefaults instantiates a new NestedModuleBay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedModuleBayWithDefaults() *NestedModuleBay {
	this := NestedModuleBay{}
	return &this
}

// GetId returns the Id field value
func (o *NestedModuleBay) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedModuleBay) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedModuleBay) SetId(v int64) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedModuleBay) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedModuleBay) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedModuleBay) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *NestedModuleBay) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *NestedModuleBay) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *NestedModuleBay) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *NestedModuleBay) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedModuleBay) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedModuleBay) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedModuleBay) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedModuleBay) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedModuleBay) SetName(v string) {
	o.Name = v
}

func (o NestedModuleBay) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedModuleBay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedModuleBay) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
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

	varNestedModuleBay := _NestedModuleBay{}

	err = json.Unmarshal(data, &varNestedModuleBay)

	if err != nil {
		return err
	}

	*o = NestedModuleBay(varNestedModuleBay)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedModuleBay struct {
	value *NestedModuleBay
	isSet bool
}

func (v NullableNestedModuleBay) Get() *NestedModuleBay {
	return v.value
}

func (v *NullableNestedModuleBay) Set(val *NestedModuleBay) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedModuleBay) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedModuleBay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedModuleBay(val *NestedModuleBay) *NullableNestedModuleBay {
	return &NullableNestedModuleBay{value: val, isSet: true}
}

func (v NullableNestedModuleBay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedModuleBay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
