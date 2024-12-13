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

// checks if the BriefModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefModuleRequest{}

// BriefModuleRequest Adds support for custom fields and tags.
type BriefModuleRequest struct {
	Device               BriefDeviceRequest     `json:"device"`
	ModuleBay            NestedModuleBayRequest `json:"module_bay"`
	AdditionalProperties map[string]interface{}
}

type _BriefModuleRequest BriefModuleRequest

// NewBriefModuleRequest instantiates a new BriefModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefModuleRequest(device BriefDeviceRequest, moduleBay NestedModuleBayRequest) *BriefModuleRequest {
	this := BriefModuleRequest{}
	this.Device = device
	this.ModuleBay = moduleBay
	return &this
}

// NewBriefModuleRequestWithDefaults instantiates a new BriefModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefModuleRequestWithDefaults() *BriefModuleRequest {
	this := BriefModuleRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *BriefModuleRequest) GetDevice() BriefDeviceRequest {
	if o == nil {
		var ret BriefDeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *BriefModuleRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *BriefModuleRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = v
}

// GetModuleBay returns the ModuleBay field value
func (o *BriefModuleRequest) GetModuleBay() NestedModuleBayRequest {
	if o == nil {
		var ret NestedModuleBayRequest
		return ret
	}

	return o.ModuleBay
}

// GetModuleBayOk returns a tuple with the ModuleBay field value
// and a boolean to check if the value has been set.
func (o *BriefModuleRequest) GetModuleBayOk() (*NestedModuleBayRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleBay, true
}

// SetModuleBay sets field value
func (o *BriefModuleRequest) SetModuleBay(v NestedModuleBayRequest) {
	o.ModuleBay = v
}

func (o BriefModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	toSerialize["module_bay"] = o.ModuleBay

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefModuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
		"module_bay",
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

	varBriefModuleRequest := _BriefModuleRequest{}

	err = json.Unmarshal(data, &varBriefModuleRequest)

	if err != nil {
		return err
	}

	*o = BriefModuleRequest(varBriefModuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module_bay")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefModuleRequest struct {
	value *BriefModuleRequest
	isSet bool
}

func (v NullableBriefModuleRequest) Get() *BriefModuleRequest {
	return v.value
}

func (v *NullableBriefModuleRequest) Set(val *BriefModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefModuleRequest(val *BriefModuleRequest) *NullableBriefModuleRequest {
	return &NullableBriefModuleRequest{value: val, isSet: true}
}

func (v NullableBriefModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
