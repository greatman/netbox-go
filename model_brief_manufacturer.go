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

// checks if the BriefManufacturer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefManufacturer{}

// BriefManufacturer Adds support for custom fields and tags.
type BriefManufacturer struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Description          *string `json:"description,omitempty"`
	DevicetypeCount      int64   `json:"devicetype_count"`
	AdditionalProperties map[string]interface{}
}

type _BriefManufacturer BriefManufacturer

// NewBriefManufacturer instantiates a new BriefManufacturer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefManufacturer(id int32, url string, display string, name string, slug string, devicetypeCount int64) *BriefManufacturer {
	this := BriefManufacturer{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.DevicetypeCount = devicetypeCount
	return &this
}

// NewBriefManufacturerWithDefaults instantiates a new BriefManufacturer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefManufacturerWithDefaults() *BriefManufacturer {
	this := BriefManufacturer{}
	return &this
}

// GetId returns the Id field value
func (o *BriefManufacturer) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefManufacturer) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefManufacturer) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefManufacturer) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefManufacturer) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefManufacturer) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefManufacturer) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefManufacturer) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefManufacturer) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefManufacturer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefManufacturer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefManufacturer) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefManufacturer) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefManufacturer) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefManufacturer) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefManufacturer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefManufacturer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefManufacturer) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefManufacturer) SetDescription(v string) {
	o.Description = &v
}

// GetDevicetypeCount returns the DevicetypeCount field value
func (o *BriefManufacturer) GetDevicetypeCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DevicetypeCount
}

// GetDevicetypeCountOk returns a tuple with the DevicetypeCount field value
// and a boolean to check if the value has been set.
func (o *BriefManufacturer) GetDevicetypeCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevicetypeCount, true
}

// SetDevicetypeCount sets field value
func (o *BriefManufacturer) SetDevicetypeCount(v int64) {
	o.DevicetypeCount = v
}

func (o BriefManufacturer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefManufacturer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["devicetype_count"] = o.DevicetypeCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefManufacturer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
		"devicetype_count",
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

	varBriefManufacturer := _BriefManufacturer{}

	err = json.Unmarshal(data, &varBriefManufacturer)

	if err != nil {
		return err
	}

	*o = BriefManufacturer(varBriefManufacturer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "devicetype_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefManufacturer struct {
	value *BriefManufacturer
	isSet bool
}

func (v NullableBriefManufacturer) Get() *BriefManufacturer {
	return v.value
}

func (v *NullableBriefManufacturer) Set(val *BriefManufacturer) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefManufacturer) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefManufacturer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefManufacturer(val *BriefManufacturer) *NullableBriefManufacturer {
	return &NullableBriefManufacturer{value: val, isSet: true}
}

func (v NullableBriefManufacturer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefManufacturer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
