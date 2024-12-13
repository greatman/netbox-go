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

// checks if the BriefVirtualCircuitRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefVirtualCircuitRequest{}

// BriefVirtualCircuitRequest Adds support for custom fields and tags.
type BriefVirtualCircuitRequest struct {
	// Unique circuit ID
	Cid                  string                      `json:"cid"`
	ProviderNetwork      BriefProviderNetworkRequest `json:"provider_network"`
	Description          *string                     `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefVirtualCircuitRequest BriefVirtualCircuitRequest

// NewBriefVirtualCircuitRequest instantiates a new BriefVirtualCircuitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefVirtualCircuitRequest(cid string, providerNetwork BriefProviderNetworkRequest) *BriefVirtualCircuitRequest {
	this := BriefVirtualCircuitRequest{}
	this.Cid = cid
	this.ProviderNetwork = providerNetwork
	return &this
}

// NewBriefVirtualCircuitRequestWithDefaults instantiates a new BriefVirtualCircuitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefVirtualCircuitRequestWithDefaults() *BriefVirtualCircuitRequest {
	this := BriefVirtualCircuitRequest{}
	return &this
}

// GetCid returns the Cid field value
func (o *BriefVirtualCircuitRequest) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualCircuitRequest) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *BriefVirtualCircuitRequest) SetCid(v string) {
	o.Cid = v
}

// GetProviderNetwork returns the ProviderNetwork field value
func (o *BriefVirtualCircuitRequest) GetProviderNetwork() BriefProviderNetworkRequest {
	if o == nil {
		var ret BriefProviderNetworkRequest
		return ret
	}

	return o.ProviderNetwork
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualCircuitRequest) GetProviderNetworkOk() (*BriefProviderNetworkRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderNetwork, true
}

// SetProviderNetwork sets field value
func (o *BriefVirtualCircuitRequest) SetProviderNetwork(v BriefProviderNetworkRequest) {
	o.ProviderNetwork = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefVirtualCircuitRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVirtualCircuitRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefVirtualCircuitRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefVirtualCircuitRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefVirtualCircuitRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefVirtualCircuitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["provider_network"] = o.ProviderNetwork
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefVirtualCircuitRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cid",
		"provider_network",
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

	varBriefVirtualCircuitRequest := _BriefVirtualCircuitRequest{}

	err = json.Unmarshal(data, &varBriefVirtualCircuitRequest)

	if err != nil {
		return err
	}

	*o = BriefVirtualCircuitRequest(varBriefVirtualCircuitRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cid")
		delete(additionalProperties, "provider_network")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefVirtualCircuitRequest struct {
	value *BriefVirtualCircuitRequest
	isSet bool
}

func (v NullableBriefVirtualCircuitRequest) Get() *BriefVirtualCircuitRequest {
	return v.value
}

func (v *NullableBriefVirtualCircuitRequest) Set(val *BriefVirtualCircuitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefVirtualCircuitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefVirtualCircuitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefVirtualCircuitRequest(val *BriefVirtualCircuitRequest) *NullableBriefVirtualCircuitRequest {
	return &NullableBriefVirtualCircuitRequest{value: val, isSet: true}
}

func (v NullableBriefVirtualCircuitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefVirtualCircuitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
