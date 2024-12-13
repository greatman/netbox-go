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

// checks if the BriefVirtualCircuit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefVirtualCircuit{}

// BriefVirtualCircuit Adds support for custom fields and tags.
type BriefVirtualCircuit struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// Unique circuit ID
	Cid                  string               `json:"cid"`
	ProviderNetwork      BriefProviderNetwork `json:"provider_network"`
	Description          *string              `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefVirtualCircuit BriefVirtualCircuit

// NewBriefVirtualCircuit instantiates a new BriefVirtualCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefVirtualCircuit(id int32, url string, display string, cid string, providerNetwork BriefProviderNetwork) *BriefVirtualCircuit {
	this := BriefVirtualCircuit{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Cid = cid
	this.ProviderNetwork = providerNetwork
	return &this
}

// NewBriefVirtualCircuitWithDefaults instantiates a new BriefVirtualCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefVirtualCircuitWithDefaults() *BriefVirtualCircuit {
	this := BriefVirtualCircuit{}
	return &this
}

// GetId returns the Id field value
func (o *BriefVirtualCircuit) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualCircuit) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefVirtualCircuit) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefVirtualCircuit) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualCircuit) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefVirtualCircuit) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefVirtualCircuit) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualCircuit) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefVirtualCircuit) SetDisplay(v string) {
	o.Display = v
}

// GetCid returns the Cid field value
func (o *BriefVirtualCircuit) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualCircuit) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *BriefVirtualCircuit) SetCid(v string) {
	o.Cid = v
}

// GetProviderNetwork returns the ProviderNetwork field value
func (o *BriefVirtualCircuit) GetProviderNetwork() BriefProviderNetwork {
	if o == nil {
		var ret BriefProviderNetwork
		return ret
	}

	return o.ProviderNetwork
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualCircuit) GetProviderNetworkOk() (*BriefProviderNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderNetwork, true
}

// SetProviderNetwork sets field value
func (o *BriefVirtualCircuit) SetProviderNetwork(v BriefProviderNetwork) {
	o.ProviderNetwork = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefVirtualCircuit) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVirtualCircuit) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefVirtualCircuit) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefVirtualCircuit) SetDescription(v string) {
	o.Description = &v
}

func (o BriefVirtualCircuit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefVirtualCircuit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
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

func (o *BriefVirtualCircuit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
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

	varBriefVirtualCircuit := _BriefVirtualCircuit{}

	err = json.Unmarshal(data, &varBriefVirtualCircuit)

	if err != nil {
		return err
	}

	*o = BriefVirtualCircuit(varBriefVirtualCircuit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "cid")
		delete(additionalProperties, "provider_network")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefVirtualCircuit struct {
	value *BriefVirtualCircuit
	isSet bool
}

func (v NullableBriefVirtualCircuit) Get() *BriefVirtualCircuit {
	return v.value
}

func (v *NullableBriefVirtualCircuit) Set(val *BriefVirtualCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefVirtualCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefVirtualCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefVirtualCircuit(val *BriefVirtualCircuit) *NullableBriefVirtualCircuit {
	return &NullableBriefVirtualCircuit{value: val, isSet: true}
}

func (v NullableBriefVirtualCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefVirtualCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
